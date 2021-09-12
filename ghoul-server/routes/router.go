package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"kidmortal.ghoul/controllers"
)

func CreateRoutes(router *mux.Router) {

<<<<<<< HEAD
	router.HandleFunc("/ws/{ChannelId}", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		controllers.HandleWebSocket(rw, r)
	}).Methods("GET")

=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
	router.HandleFunc("/bot", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		controllers.HandleGetBot(rw, r)
	}).Methods("GET")

	router.HandleFunc("/bot", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		controllers.HandleBotLogin(rw, r)
	}).Methods("POST")

	router.HandleFunc("/guilds", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		controllers.HandleBotGetGuilds(rw, r)
	}).Methods("GET")

<<<<<<< HEAD
	router.HandleFunc("/messages/{ChannelId}", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		controllers.HandleGetMessages(rw, r)
	}).Methods("GET")

=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
}
