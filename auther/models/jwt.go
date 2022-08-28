package models

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtSigner wraps the signing key and the issuer
type JwtSigner struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
}

type JwtClaim struct {
	Name, Email string
	IsAdmin     bool
	jwt.StandardClaims
}

var (
	ErrorUnparsableClaims = errors.New("couldn't parse JWT claims")
	ErrorExpiredToken     = errors.New("JWT token is expired")
)

func NewJwtSigner(secretKey, issuerName string, expirationMinutes int64) JwtSigner {
	return JwtSigner{secretKey, issuerName, expirationMinutes}
}

// GenerateToken generates a jwt token
func (j *JwtSigner) GenerateToken(user User) (signedToken string, err error) {
	claims := &JwtClaim{
		Email:   user.Email,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))

	return
}

// ValidateToken validates the jwt token
// if error is nil, the token is valid
func (j *JwtSigner) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = ErrorUnparsableClaims
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = ErrorExpiredToken
		return
	}

	return
}
