package routes

import (
	"github.com/diogo-costa98/deck-of-cards/api/handlers"
	"github.com/gorilla/mux"
)

func SetDeckRoutes(router *mux.Router) {
	router.HandleFunc("/deck", handlers.CreateDeckHandler).Methods("POST")
	router.HandleFunc("/deck/{id}", handlers.OpenDeckHandler).Methods("GET")
	router.HandleFunc("/deck/{id}/draw", handlers.DrawCardHandler).Methods("POST")
}
