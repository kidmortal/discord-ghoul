package controllers

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
<<<<<<< HEAD
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	"kidmortal.ghoul/services"
	"kidmortal.ghoul/utils"
)

<<<<<<< HEAD
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
=======
var contador = 0
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b

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
<<<<<<< HEAD
	botUser, err := services.LoginDiscord(params.Secret)
	if err != nil {
		SendResponse(err.Error(), rw)
		return
	}
=======
	botUser := services.LoginDiscord(params.Secret)
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	jsonString, err := json.Marshal(botUser)
	if err != nil {
		log.Fatal("Erro marshal bot")
	}
<<<<<<< HEAD
	fmt.Println("Success login")
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
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
<<<<<<< HEAD
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
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	rw.Write(j)
}

func HandleGetBot(rw http.ResponseWriter, r *http.Request) {
	bot, err := services.GetBot()
	if err != nil {
		SendResponse(err.Error(), rw)
		return
	}
<<<<<<< HEAD
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
=======
	rw.Write(StrucToByte(bot))

>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
}
