package handlers

import (
	"auther/models"
	"net/http"
)

type KeyUser struct{}

func (d *backpack) HandleUserSignup(rw http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value(KeyUser{}).(models.User)
	usr.IsAdmin = false
	d.Logger.Info("Signin up user", usr)
	err := d.CreateUser(usr)

	if err != nil {
		d.Logger.Error("Error creating the user", err)
		http.Error(rw, "Error during the signup process", http.StatusBadRequest)
	}
	// 204
	rw.WriteHeader(http.StatusCreated)
}
