package services

import (
<<<<<<< HEAD
	"encoding/json"
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
<<<<<<< HEAD
	"github.com/gorilla/websocket"
)

type SocketSendMessage struct {
	Call    string      `json:"call"`
	Payload interface{} `json:"payload"`
}

var DiscordSession *discordgo.Session
var SocketConn *websocket.Conn
var SelectedChannel string

func LoginDiscord(key string) (*discordgo.User, error) {
	if DiscordSession != nil {
		fmt.Println("Already connected")
		return DiscordSession.State.User, nil
=======
)

var DiscordSession *discordgo.Session

func LoginDiscord(key string) *discordgo.User {

	if DiscordSession != nil {
		fmt.Println("Already connected")
		return DiscordSession.State.User
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	}

	discord, err := discordgo.New("Bot " + key)
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error opening connection")
	}
	fmt.Println("Connection Success!")
	DiscordSession = discord
	return discord.State.User, nil
=======
	discord.AddHandler(messageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return nil
	}
	fmt.Println("Connection Success!")
	DiscordSession = discord
	return discord.State.User
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

<<<<<<< HEAD
	if m.ChannelID == SelectedChannel {
		WriteSocketMessage("NEW_MESSAGES", m.Message)
	}
}

func GetAllGuilds() ([]*discordgo.Guild, error) {
	if DiscordSession == nil {
		return nil, errors.New("bot is not logged in")
	}
	return DiscordSession.State.Guilds, nil
}

func GetChannelMessages(id string) ([]*discordgo.Message, error) {
	if DiscordSession == nil {
		return nil, errors.New("bot is not logged in")
	}
	messages, err := DiscordSession.ChannelMessages(id, 30, "", "", "")
	if err != nil {
		log.Fatal(err)
	}
	return messages, nil
=======
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
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
}

func GetBot() (*discordgo.User, error) {
	if DiscordSession == nil {
		return nil, errors.New("bot is not logged in")
	}
	return DiscordSession.State.User, nil
}
<<<<<<< HEAD

func WebSocketChannelSession(conn *websocket.Conn, channelId string) {
	if DiscordSession == nil {
		log.Println("Bot not connected")
		return
	}
	messages, err := DiscordSession.ChannelMessages(channelId, 30, "", "", "")
	if err != nil {
		log.Fatal(err)
	}
	SocketConn = conn
	SelectedChannel = channelId
	WriteSocketMessage("NEW_MESSAGES", messages)
	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		DiscordSession.ChannelMessageSend(channelId, string(p))

	}
}

func WriteSocketMessage(call string, payload interface{}) {
	message := SocketSendMessage{Call: call, Payload: payload}
	j, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	if err := SocketConn.WriteMessage(1, j); err != nil {
		log.Println(err)
	}
}
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
