package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var DiscordSession *discordgo.Session

func LoginDiscord(key string) *discordgo.User {

	if DiscordSession != nil {
		fmt.Println("Already connected")
		return DiscordSession.State.User
	}

	discord, err := discordgo.New("Bot " + key)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(messageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return nil
	}
	fmt.Println("Connection Success!")
	DiscordSession = discord
	return discord.State.User
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		fmt.Println(s.State.Guilds)
		for _, v := range s.State.Guilds {
			fmt.Println(v.Name)
		}
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func GetAllGuilds() ([]*discordgo.Guild, error) {
	if DiscordSession == nil {
		return nil, errors.New("bot is not logged in")
	}
	return DiscordSession.State.Guilds, nil
}

func GetBot() (*discordgo.User, error) {
	if DiscordSession == nil {
		return nil, errors.New("bot is not logged in")
	}
	return DiscordSession.State.User, nil
}
