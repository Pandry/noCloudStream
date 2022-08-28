package handlers

import (
	"auther/models"
	"html"
	"net/http"
)

func (d backpack) SanitizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		providedCredentials := r.Context().Value(KeyUser{}).(models.User)

		providedCredentials.Name = html.EscapeString(providedCredentials.Name)
		providedCredentials.Email = html.EscapeString(providedCredentials.Email)
		providedCredentials.Password = html.EscapeString(providedCredentials.Password)

		next.ServeHTTP(rw, r)
	})
}
