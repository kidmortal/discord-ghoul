package controllers

import (
	"encoding/json"
	"log"
	"net/http"

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

func HandleBotLogin(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params loginParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}
	session := services.LoginDiscord(params.Secret)
	jsonString, err := json.Marshal(session.State.User)
	if err != nil {
		log.Fatal("Erro marshal bot")
	}
	rw.Write(jsonString)
}
