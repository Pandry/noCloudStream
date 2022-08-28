// Package classification auther API.
//
// Documentation for the auther microservice
//
//	Schemes: http,https
//	BasePath: /api/auth
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security
//	- jwt:
//
//	SecurityDefinitions:
//	jwt:
//		type: JWT
//		name: jtw
//		in: cookies
//
// swagger:meta

package handlers

import (
	"auther/models"
	"net/http"
)

const CookieName = "Session"

func (d *backpack) HandleUserLogin(rw http.ResponseWriter, r *http.Request) {

	providedCredentials := r.Context().Value(KeyUser{}).(models.User)
	user, err := d.Repository.GetUserByMail(providedCredentials.Email)
	if err != nil {
		d.Logger.Error("Error fetching user", err)
		http.Error(rw, "Error during the login process. Is the user signed up?", http.StatusBadRequest)
		return
	}
	loggedIn := models.ComparePassword(user.Password, providedCredentials.Password)
	providedCredentials.Password = ""
	if !loggedIn {
		d.Logger.Error("Failed login for user", "user", providedCredentials.Email)
		http.Error(rw, "Wrong password", http.StatusUnauthorized)
		return
	}
	token, err := d.Signer.GenerateToken(user)
	if err != nil {
		d.Logger.Error("Error generating the token", err)
		http.Error(rw, "Error during the login process.", http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Write([]byte(`{"jtw": "` + token + `"}`))

}
