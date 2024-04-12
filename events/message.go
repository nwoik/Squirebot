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

	channel, _ := session.Channel(message.ChannelID)
	guild, _ := session.Guild(message.GuildID)

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

		if command == "guildinfo" {
			embed := NewRichEmbed(guild.Name, "Server Info", 0x00bfff)

			embed.SetThumbnail(guild.IconURL(""))

			// embed.AddField("**Owner:**", , false)
			embed.AddField("**Member Count:**", fmt.Sprint(len(guild.Members)), false)
			embed.AddField("**Role Count:**", fmt.Sprint(len(guild.Roles)), false)
			embed.AddField("**Booster Count:**", fmt.Sprint(guild.PremiumSubscriptionCount), false)
			embed.SetFooter(fmt.Sprintf("Requested by %s", message.Author.Username), message.Author.AvatarURL(""))

			session.ChannelMessageSendEmbed(channel.ID, embed.MessageEmbed)
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
			log.Println(channel.Members)
			log.Println(GetUserByID(guild, guild.OwnerID))
			return
		}
	}
}

func GetUserByID(guild *discordgo.Guild, userID string) *discordgo.Member {
	log.Println(guild.Members)
	for _, v := range guild.Members {
		log.Println(v)
		// if v.User.ID == userID {
		// 	return v
		// }
	}
	return nil
}
