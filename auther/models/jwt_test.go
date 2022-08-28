package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateValidateToken(t *testing.T) {
	jwtWrapper := NewJwtSigner("verysecretkey", "AuthService", 10)

	userPassword := "mySecretPass"

	user := User{
		Name:     "Name",
		Email:    "me@example.com",
		Password: userPassword,
	}
	_, err := user.HashUserPassword()
	assert.NoError(t, err)

	generatedToken, err := jwtWrapper.GenerateToken(user)
	assert.NoError(t, err)

	claims, err := jwtWrapper.ValidateToken(generatedToken)
	assert.NoError(t, err)

	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, jwtWrapper.Issuer, claims.Issuer)
}
