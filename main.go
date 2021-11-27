package main

import (
	// "github.com/feifeijun97/GenshinBackEnd/repository"

	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/feifeijun97/GenshinBackEnd/modules/character"
	"github.com/feifeijun97/GenshinBackEnd/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//handle the routes for API request
func apiRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}

func main() {
	repository.ConnectToPostgreDb()
	c := character.Character{}

	c.GetCharacterById(1)
	//listen to API request from client
	router := apiRouter()

	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), router))
}
