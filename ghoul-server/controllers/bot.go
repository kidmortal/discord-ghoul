package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"kidmortal.ghoul/services"
	"kidmortal.ghoul/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type loginParams struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}

func SendResponse(response string, rw http.ResponseWriter) {
	jsonString := utils.StructToJsonString(map[string]string{"response": response})
	rw.Write([]byte(jsonString))
}

func StrucToByte(struc interface{}) []byte {
	j, err := json.Marshal(struc)
	if err != nil {
		log.Fatal("Error sendjson")
	}
	return j
}

func HandleBotLogin(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params loginParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}
	botUser, err := services.LoginDiscord(params.Secret)
	if err != nil {
		SendResponse(err.Error(), rw)
		return
	}
	jsonString, err := json.Marshal(botUser)
	if err != nil {
		log.Fatal("Erro marshal bot")
	}
	fmt.Println("Success login")
	rw.Write(jsonString)
}

func HandleBotGetGuilds(rw http.ResponseWriter, r *http.Request) {
	guilds, err := services.GetAllGuilds()
	if err != nil {
		SendResponse(err.Error(), rw)
		return
	}
	var guildsArr []*discordgo.Guild
	for _, v := range guilds {
		guildsArr = append(guildsArr, v)
	}
	j, err := json.Marshal(guildsArr)
	if err != nil {
		log.Fatal("marshal error")
	}
	fmt.Println("Success get guilds")
	rw.Write(j)
}

func HandleGetMessages(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["ChannelId"]
	fmt.Println(id)
	messages, err := services.GetChannelMessages(id)
	if err != nil {
		log.Fatal(err)
	}
	var messagesArr []*discordgo.Message
	for _, v := range messages {
		messagesArr = append(messagesArr, v)
	}
	j, err := json.Marshal(messagesArr)
	if err != nil {
		log.Fatal("marshal error")
	}
	fmt.Println("Success get messages")
	rw.Write(j)
}

func HandleGetBot(rw http.ResponseWriter, r *http.Request) {
	bot, err := services.GetBot()
	if err != nil {
		SendResponse(err.Error(), rw)
		return
	}
	fmt.Println("Success get bot")
	rw.Write(StrucToByte(bot))
}

func HandleWebSocket(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["ChannelId"]

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sucess websocket")
	services.WebSocketChannelSession(ws, id)
}
