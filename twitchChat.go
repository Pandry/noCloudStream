package main

import (
	"gopkg.in/irc.v3"
	"log"
	"net"
	"strings"
)

type IrcChatMessage struct {
	Author string
	MessageContent string
	ID string
}

func syncChatLoop(channelName string, ch chan IrcChatMessage) {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatalln(err)
	}

	config := irc.ClientConfig{
		Nick: "justinfan854",
		//Pass: "password",
		User: "justinfan854",
		Name: "noCloudStream Chat Crawler",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				c.Write("CAP REQ :twitch.tv/tags twitch.tv/commands")
				// 001 is a welcome event, so we join channels there
				c.Write("JOIN #" + channelName)
			} else /*if m.Command == "PRIVMSG" && c.FromChannel(m) {
				// Create a handler on all messages.
				c.WriteMessage(&irc.Message{
					Command: "PRIVMSG",
					Params: []string{
						m.Params[0],
						m.Trailing(),
					},
				})
			}*/if m.Command == "PRIVMSG" && c.FromChannel(m) {
				if debug {for k,v := range m.Tags {
					log.Printf(k + ": " + v.Encode())
				}}
				aName := m.Prefix.Name;
				if tName, ok:= m.Tags.GetTag("display-name"); ok {
					aName = tName
				}
				aID, _ := m.Tags.GetTag("user-id")
				ch <- IrcChatMessage{
					//Author:         ,
					Author:         aName,
					ID: aID,
					MessageContent: strings.Join(m.Params[1:], " "),
				}
				if debug {log.Printf(m.Prefix.Name + ": " + strings.Join(m.Params[1:], " "))}
			}
		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("IRC crashed lol")
}
