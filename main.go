package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"html"
	"net/http"
)

//Changelog
//Cleanup
//removed hardcoded channel name (now flag)
//Readme
//Log lib

var addr string
var channelName = ""

func main() {
	debug := flag.Bool("debug", false, "Enables detailed output.")
	trace := flag.Bool("trace", false, "Enables even more detailed output.")
	flag.StringVar(&channelName, "channel", "meetscience", "The Twitch channel name to attach to")
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "http service address")
	flag.Parse()
	log.WithField("channelName", channelName).WithField("address", addr).Info("Starting noCloudStream")

	if *debug {
		log.SetLevel(log.DebugLevel)
	}
	if *trace {
		log.SetLevel(log.TraceLevel)
	}

	log.Trace("Instantiating the chat hub")
	hub := newHub()
	log.Trace("Starting the chat hub")
	go hub.run()

	log.Trace("Creating the twitch message queue")
	twichChat := make(chan IrcChatMessage, 100)
	log.Trace("Starting the IRC client")
	go syncChatLoop(channelName, twichChat)
	log.Trace("Starting the IRC to chat loop")
	go func() {
		for {
			msg := <-twichChat
			log.WithField("message", msg).Trace("Sending message from the IRC channel to the WS chan")
			//Yes, I know hardcoding stuff sucks. But I also do @ coding, so yeah. It's a side project.
			hub.broadcast <- []byte(`{"topic":"chat","authorId": ` + html.EscapeString(msg.ID) +
				`, "author": "` + html.EscapeString(msg.Author) +
				`","message":"` + html.EscapeString(msg.MessageContent) + `"}`)
		}
	}()

	//Routes funcions (see the httpRoutes.go file)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/obs", serveObs)
	http.HandleFunc("/assets.js", serveJs)
	http.HandleFunc("/ws", hub.serveWs)

	listenAddr := addr
	if listenAddr[0] == ':' {
		listenAddr = "0.0.0.0"
		log.Warn("Hey, you're listening on the whole nic, watch out!")
	}
	log.WithField("url", "http://"+addr).Info("Started serving")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
