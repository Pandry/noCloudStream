package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/irc.v3"
	"net"
	"strings"
)

//IrcChatMessage is the struct we use to pass the message from the
//IRC channel to the ws chat (via the chan)
type IrcChatMessage struct {
	Author         string
	MessageContent string
	ID             string
}

func syncChatLoop(channelName string, ch chan IrcChatMessage) {
	ircAddress := "irc.chat.twitch.tv:6667"
	log.WithField("endpoint", ircAddress).Trace("Connecting to the twitch IRC endpoint")
	conn, err := net.Dial("tcp", ircAddress)
	if err != nil {
		log.Fatalln(err)
	}
	log.WithField("endpoint", ircAddress).Info("Connecting to the twitch IRC endpoint")

	//TODO (Pandry): ATM an hardcoded nick should be just fine
	//					Also Twitch is ok with multiple users with the same nick/users
	//					AFA I was able to read online.
	//					Eventually an API integration could be ok but still... It's public
	//					stuff, I don't like I need an API key to gather such data(even if I
	//					undertand the reasons)

	config := irc.ClientConfig{
		Nick: "justinfan854",
		//Pass: "password",
		User: "justinfan854",
		Name: "noCloudStream Chat Crawler",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				log.WithField("channelName", "#"+channelName).Debug("Sending the CAP and JOIN message to the IRC server")
				c.Write("CAP REQ :twitch.tv/tags twitch.tv/commands")
				c.Write("JOIN #" + channelName)
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {

				log.WithFields(log.Fields{
					"tags":    m.Tags,
					"message": m.Params,
					"prefix":  m.Prefix,
				}).Debug("Received a message from the server")

				aName := m.Prefix.Name
				if tName, ok := m.Tags.GetTag("display-name"); ok {
					aName = tName
				}
				aID, _ := m.Tags.GetTag("user-id")
				ch <- IrcChatMessage{
					Author:         aName,
					ID:             aID,
					MessageContent: strings.Join(m.Params[1:], " "),
				}
			}
		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatal("Yeah. Life sucks and shit happens. And software sucks too. We've got out of a server somehow. Please open a PR.")
}
