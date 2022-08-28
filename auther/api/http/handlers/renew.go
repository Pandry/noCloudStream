package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func (d *backpack) HandleRenew(rw http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		http.Error(rw, "Unable to get the session", http.StatusBadRequest)
	}

	token := c.Value
	claims, err := d.Signer.ValidateToken(token)
	if err != nil {
		http.Error(rw, "Invalid token", http.StatusBadRequest)
	}

	user, err := d.Repository.GetUserByMail(claims.Email)
	if err != nil {
		d.Logger.Error("Error fetching user", err)
		http.Error(rw, "Error during the renew process. Is the user signed up?", http.StatusBadRequest)
	}
	newToken, err := d.Signer.GenerateToken(user)
	if err != nil {
		d.Logger.Error("Error generating the token", err)
		http.Error(rw, "Error during the login process.", http.StatusInternalServerError)
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Write([]byte(`{"jtw": "` + newToken + `"}`))

}

func (d *backpack) HandleInfo(rw http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		http.Error(rw, "Unable to get the session", http.StatusBadRequest)
	}

	token := c.Value
	claims, err := d.Signer.ValidateToken(token)
	if err != nil {
		http.Error(rw, "Invalid token", http.StatusBadRequest)
	}
	expiry := time.Unix(claims.ExpiresAt, 0)
	rw.Write([]byte(fmt.Sprintf("Name: %s, Expiry: %s, Email: %s", claims.Name, expiry, claims.Email)))
}
