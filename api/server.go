package api

import (
	"fmt"
	"net/http"

	"github.com/diogo-costa98/deck-of-cards/api/routes"
	"github.com/gorilla/mux"
)

func StartServer() error {
	router := mux.NewRouter()

	// Define "/deck" the HTTP request handlers
	routes.SetDeckRoutes(router)

	// Start the HTTP server
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}

	return nil
}
