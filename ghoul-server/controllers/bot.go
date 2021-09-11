package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"kidmortal.ghoul/services"
	"kidmortal.ghoul/utils"
)

var contador = 0

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
	botUser := services.LoginDiscord(params.Secret)
	jsonString, err := json.Marshal(botUser)
	if err != nil {
		log.Fatal("Erro marshal bot")
	}
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
	rw.Write(j)
}

func HandleGetBot(rw http.ResponseWriter, r *http.Request) {
	bot, err := services.GetBot()
	if err != nil {
		SendResponse(err.Error(), rw)
		return
	}
	rw.Write(StrucToByte(bot))

}
