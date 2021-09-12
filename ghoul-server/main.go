package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"kidmortal.ghoul/routes"
)

func main() {
	router := mux.NewRouter()
	routes.CreateRoutes(router)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	fmt.Println("🎉 Running GO server on port 5000")
	log.Fatal((http.ListenAndServe(":5000", c.Handler(router))))
}
