package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Deck represents a deck of playing cards.
type Deck struct {
	ID        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []Card    `json:"cards,omitempty"`
}

// Creates a new deck shuffled/unshuffled with the selected playing cards
func CreateDeck(cards []string, shuffled bool) *Deck {
	createdCards := CreateCards(cards)

	deck := &Deck{
		ID:        uuid.New(),
		Shuffled:  shuffled,
		Remaining: len(createdCards),
		Cards:     createdCards,
	}

	if deck.Shuffled {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(deck.Cards), func(i, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}

	return deck
}

// Draws a certain amount of playing cards from a deck
func DrawDeck(count int, deck *Deck) ([]Card, bool) {
	if count > deck.Remaining || count < 1 {
		return nil, false
	}

	drawnCards := make([]Card, count)
	for i := 0; i < count; i++ {
		drawnCards[i] = deck.Cards[deck.Remaining-i-1]
	}

	return drawnCards, true
}
