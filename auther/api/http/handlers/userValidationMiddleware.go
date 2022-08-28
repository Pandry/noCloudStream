package handlers

import (
	"auther/models"
	"context"
	"net/http"
)

func (d backpack) ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := user.FromJSON(r.Body)
		if err != nil {
			d.Logger.Error("Cannot decode user body", err)
			http.Error(rw, "Error reading the request body", http.StatusBadRequest)
		}
		if err = user.Validate(); err != nil {
			d.Logger.Error("Cannot validate user body", err)
			http.Error(rw, "Error reading the request body", http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
