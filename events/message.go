package events

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "£"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.SessionID {
		return
	}

	channel, err := session.Channel(message.ChannelID)

	if err != nil {
		log.Fatal(err)
	}

	args := strings.Split(message.Content, " ")
	if strings.Contains(args[0], prefix) {
		fmt.Println(args)
		command := strings.Trim(args[0], prefix)

		if command == "ban" {
			for _, member := range message.Mentions {
				re, _ := regexp.Compile(`(\W\w+\s(.+\d+\W+)+)`)
				reason := re.ReplaceAllString(message.Content, "")
				session.GuildBanCreateWithReason(channel.GuildID, member.ID, reason, 2)
			}
		}

		if command == "kick" {
			for _, member := range message.Mentions {
				re, _ := regexp.Compile(`(\W\w+\s(.+\d+\W+)+)`)
				reason := re.ReplaceAllString(message.Content, "")
				session.GuildMemberDeleteWithReason(channel.GuildID, member.ID, reason)
			}
		}

		if command == "ping" {
			session.ChannelMessageSend(message.ChannelID, "pong")
			return
		}

	}
}
