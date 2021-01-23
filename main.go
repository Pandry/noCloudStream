package main

import (
	"flag"
	"log"
	"net/http"
	"html"
)

var addr = flag.String("addr", ":8080", "http service address")
var debug = false

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()

	twichChat := make(chan IrcChatMessage, 100)
	go syncChatLoop("zizaran", twichChat)
	go func(){
		for {
		msg := <- twichChat
		hub.broadcast <- []byte(`{"topic":"chat","authorId": `+ html.EscapeString(msg.ID) +`, "author": "`+html.EscapeString(msg.Author)+`","message":"`+html.EscapeString(msg.MessageContent)+`"}`)
		
	}
	}()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/obs", serveObs)
	http.HandleFunc("/assets.js", func(w http.ResponseWriter, r *http.Request) {w.Header().Add("Content-Type", "text/javascript"); http.ServeFile(w, r, "anime.min.js") })
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
