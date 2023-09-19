package events

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "£"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// fmt.Println(message.Author.Username, message.Content)

	if message.Author.ID == session.State.SessionID {
		return
	}

	args := strings.Split(message.Content, " ")
	fmt.Println(args)
	if strings.Contains(args[0], prefix) {
		command := strings.Trim(args[0], prefix)

		if command == "ping" {
			session.ChannelMessageSend(message.ChannelID, "pong")
		}

	}
}
