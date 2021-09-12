package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
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
	}

	discord, err := discordgo.New("Bot " + key)
	if err != nil {
		log.Fatal(err)
	}

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
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

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
}

func GetBot() (*discordgo.User, error) {
	if DiscordSession == nil {
		return nil, errors.New("bot is not logged in")
	}
	return DiscordSession.State.User, nil
}

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
