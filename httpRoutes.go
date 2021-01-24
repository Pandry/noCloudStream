package main

import (
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "assets/home.html")
}

func serveJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")
	http.ServeFile(w, r, "assets/js/anime.min.js")
}

func serveObs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/obs.html")
}

func (hub *Hub) serveWs(w http.ResponseWriter, r *http.Request) {
	serveWs(hub, w, r)
}
