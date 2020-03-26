package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic("unable to create discord session")
	}

	defer dg.Close()

	dg.AddHandler(messageCreate)

	if err := dg.Open(); err != nil {
		panic("unable to start bot")
	}

	fmt.Println("flopbot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, "Hello")
	if err != nil {
		fmt.Printf("error sending message: %s", err)
	}
}
