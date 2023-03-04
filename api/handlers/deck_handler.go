package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/diogo-costa98/deck-of-cards/api/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Non-persistent deck to simulate saving to a databse
var decks = make(map[uuid.UUID]*models.Deck)

// CreateDeckHandler handles the POST /deck request to create a new deck of cards.
func CreateDeckHandler(w http.ResponseWriter, r *http.Request) {
	//Handle parameters
	shuffled := r.FormValue("shuffle") == "true"

	cards := strings.Split(r.FormValue("cards"), ",")
	if len(cards) == 1 && cards[0] == "" {
		cards = []string{}
	}

	//Create new deck
	deck := models.CreateDeck(cards, shuffled)

	//Save to the (non-persistent) database
	decks[deck.ID] = deck

	//Prepare response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&models.Deck{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	})
}

// OpenDeckHandler handles the GET /deck request to get a deck of cards.
func OpenDeckHandler(w http.ResponseWriter, r *http.Request) {
	//Handle parameters
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return
	}

	//Get from non-persistent database
	deck, ok := decks[id]
	if !ok {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}

	//Prepare response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}

func DrawCardHandler(w http.ResponseWriter, r *http.Request) {
	//Handle parameters
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid deck ID", http.StatusBadRequest)
		return
	}

	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		return
	}

	//Get from non-persistent database
	deck, ok := decks[id]
	if !ok {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}

	//Draw from deck
	drawnCards, ok := models.DrawDeck(count, deck)
	if !ok {
		http.Error(w, "Invalid count parameter", http.StatusBadRequest)
		return
	}

	//Update non-persistent database
	deck.Remaining -= count
	deck.Cards = deck.Cards[:deck.Remaining]

	//Prepare response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Cards []models.Card `json:"cards"`
	}{
		Cards: drawnCards,
	})
}
