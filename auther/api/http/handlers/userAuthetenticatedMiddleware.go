package handlers

import (
	"auther/models"
	"net/http"
	"strings"
)

func (d backpack) NeedsAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Authorization")
		splitToken := strings.Split(jwtToken, "Bearer ")
		if len(splitToken) != 2 {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		jwtToken = splitToken[1]

		claims, err := d.Signer.ValidateToken(jwtToken)
		if err != nil {
			if err == models.ErrorExpiredToken {
				d.Logger.Debug("User tried to login with expired token", "user", claims.Email)
			}
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
