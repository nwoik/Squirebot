package events

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "£"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	args := strings.Split(message.Content, " ")
	fmt.Println(message.ContentWithMentionsReplaced())

	if message.Author.ID == session.State.SessionID {
		return
	}

	if strings.Index(prefix, args[0]) == 0 {
		session.ChannelMessageSend(message.ID, "This is a command response")
	}
}
