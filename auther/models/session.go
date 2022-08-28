package models

import "time"

type Session struct {
	User      string    `json:"email"`
	IP        string    `json:"ip"`
	LoginTime time.Time `json:"login_time"`
}

func AddSession(user, IP string) {
	//session := Session{user, IP, time.Now()}

}
