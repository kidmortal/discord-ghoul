package services

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func LoginDiscord(key string) *discordgo.Session {
	discord, err := discordgo.New("Bot " + key)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(messageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return discord
	}
	fmt.Println("Success")
	return discord
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
