package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"kidmortal.ghoul/controllers"
)

func CreateRoutes(router *mux.Router) {

	router.HandleFunc("/bot", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		controllers.HandleBotLogin(rw, r)
	}).Methods(http.MethodPost)

}
