package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "£"

func main() {
	token := os.Getenv("SQUIRE_TOKEN")
	fmt.Println(token)
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(CommandHandler)
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

func CommandHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	args := strings.Split(message.Content, " ")

	if message.Author.ID == session.State.SessionID {
		return
	}

	if args[0] != prefix {
		return
	}
}
