package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("SQUIRE_TOKEN")
	fmt.Println(token)
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatal(err)
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()
	fmt.Println("Ready to serve... ")
	session.UpdateGameStatus(1, "Running on Golang")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// func WelcomeHandler(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
// 	guild, err := session.Guild(event.GuildID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
