package tdd

import (
	"testing"

	"github.com/diogo-costa98/deck-of-cards/api/models"
)

func TestCreateDeck(t *testing.T) {

	testCases := []struct {
		name             string
		cardCodes        []string
		orderedCardCodes []string
		shuffled         bool
		expectedCards    int
		expectedShuf     bool
	}{
		{
			name:      "test unshuffled deck with no cards specified",
			cardCodes: nil,
			orderedCardCodes: []string{
				"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
				"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
				"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
				"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
			},
			shuffled:      false,
			expectedCards: 52,
			expectedShuf:  false,
		},
		{
			name:      "test shuffled deck with no cards specified",
			cardCodes: nil,
			orderedCardCodes: []string{
				"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
				"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
				"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
				"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
			},
			shuffled:      true,
			expectedCards: 52,
			expectedShuf:  true,
		},
		{
			name:      "test unshuffled deck with partial card codes specified",
			cardCodes: []string{"AS", "AC", "KD", "2C", "KH"},
			orderedCardCodes: []string{
				"AS", "KD", "AC", "2C", "KH",
			},
			shuffled:      false,
			expectedCards: 5,
			expectedShuf:  false,
		},
		{
			name:      "test shuffled deck with partial card codes specified",
			cardCodes: []string{"AS", "AC", "KD", "2C", "KH"},
			orderedCardCodes: []string{
				"AS", "KD", "AC", "2C", "KH",
			},
			shuffled:      true,
			expectedCards: 5,
			expectedShuf:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			deck := models.CreateDeck(tc.cardCodes, tc.shuffled)
			if len(deck.Cards) != tc.expectedCards {
				t.Errorf("Expected %d cards in deck, but got %d", tc.expectedCards, len(deck.Cards))
			}
			if deck.Shuffled != tc.expectedShuf {
				if tc.expectedShuf {
					t.Errorf("Expected deck to be shuffled, but it is unshuffled")
				} else {
					t.Errorf("Expected deck to be unshuffled, but it is shuffled")
				}
			}
			if deck.Remaining != len(deck.Cards) {
				t.Errorf("Deck.Remaining differs from its' card number")
			}

			cardsUnshuffled := true
			// Check that deck.Cards are in sequential order if deck is unshuffled
			for i := 0; i < len(deck.Cards); i++ {
				expectedCode := tc.orderedCardCodes[i]
				if deck.Cards[i].Code != expectedCode {
					cardsUnshuffled = false
					break
				}
			}

			if cardsUnshuffled && deck.Shuffled {
				t.Errorf("Deck defined as shuffled but its cards are unshuffled")
			}

			if !cardsUnshuffled && !deck.Shuffled {
				t.Errorf("Deck defined as unshuffled but its cards are shuffled")
			}
		})
	}
}

func TestDrawDeck(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		initialDeck    *models.Deck
		numCardsToDraw int
		remainingCards int
		expectedCards  []models.Card
		expectedOk     bool
	}{
		{
			name:           "Draw two cards from unshuffled deck",
			initialDeck:    models.CreateDeck(nil, false),
			numCardsToDraw: 2,
			expectedCards: []models.Card{
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
				{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
			},
			expectedOk: true,
		},
		{
			name:           "Draw all cards from unshuffled deck",
			initialDeck:    models.CreateDeck(nil, false),
			numCardsToDraw: 52,
			expectedCards: []models.Card{
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
				{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
				{Value: "JACK", Suit: "HEARTS", Code: "JH"},
				{Value: "10", Suit: "HEARTS", Code: "10H"},
				{Value: "9", Suit: "HEARTS", Code: "9H"},
				{Value: "8", Suit: "HEARTS", Code: "8H"},
				{Value: "7", Suit: "HEARTS", Code: "7H"},
				{Value: "6", Suit: "HEARTS", Code: "6H"},
				{Value: "5", Suit: "HEARTS", Code: "5H"},
				{Value: "4", Suit: "HEARTS", Code: "4H"},
				{Value: "3", Suit: "HEARTS", Code: "3H"},
				{Value: "2", Suit: "HEARTS", Code: "2H"},
				{Value: "ACE", Suit: "HEARTS", Code: "AH"},
				{Value: "KING", Suit: "CLUBS", Code: "KC"},
				{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
				{Value: "JACK", Suit: "CLUBS", Code: "JC"},
				{Value: "10", Suit: "CLUBS", Code: "10C"},
				{Value: "9", Suit: "CLUBS", Code: "9C"},
				{Value: "8", Suit: "CLUBS", Code: "8C"},
				{Value: "7", Suit: "CLUBS", Code: "7C"},
				{Value: "6", Suit: "CLUBS", Code: "6C"},
				{Value: "5", Suit: "CLUBS", Code: "5C"},
				{Value: "4", Suit: "CLUBS", Code: "4C"},
				{Value: "3", Suit: "CLUBS", Code: "3C"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "ACE", Suit: "CLUBS", Code: "AC"},
				{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
				{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
				{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
				{Value: "10", Suit: "DIAMONDS", Code: "10D"},
				{Value: "9", Suit: "DIAMONDS", Code: "9D"},
				{Value: "8", Suit: "DIAMONDS", Code: "8D"},
				{Value: "7", Suit: "DIAMONDS", Code: "7D"},
				{Value: "6", Suit: "DIAMONDS", Code: "6D"},
				{Value: "5", Suit: "DIAMONDS", Code: "5D"},
				{Value: "4", Suit: "DIAMONDS", Code: "4D"},
				{Value: "3", Suit: "DIAMONDS", Code: "3D"},
				{Value: "2", Suit: "DIAMONDS", Code: "2D"},
				{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
				{Value: "KING", Suit: "SPADES", Code: "KS"},
				{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
				{Value: "JACK", Suit: "SPADES", Code: "JS"},
				{Value: "10", Suit: "SPADES", Code: "10S"},
				{Value: "9", Suit: "SPADES", Code: "9S"},
				{Value: "8", Suit: "SPADES", Code: "8S"},
				{Value: "7", Suit: "SPADES", Code: "7S"},
				{Value: "6", Suit: "SPADES", Code: "6S"},
				{Value: "5", Suit: "SPADES", Code: "5S"},
				{Value: "4", Suit: "SPADES", Code: "4S"},
				{Value: "3", Suit: "SPADES", Code: "3S"},
				{Value: "2", Suit: "SPADES", Code: "2S"},
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
			},
			expectedOk: true,
		},
		{
			name:           "Draw more cards than remaining in the deck",
			initialDeck:    models.CreateDeck(nil, false),
			numCardsToDraw: 53,
			expectedCards:  []models.Card{},
			expectedOk:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			drawnCards, err := models.DrawDeck(tc.numCardsToDraw, tc.initialDeck)

			if err != tc.expectedOk {
				t.Errorf("Expected ok = %t, but got ok = %t", tc.expectedOk, err)
			}

			if len(drawnCards) != len(tc.expectedCards) {
				t.Errorf("Expected to draw %d cards, but drew %d", len(tc.expectedCards), len(drawnCards))
			}

			for i, expectedCard := range tc.expectedCards {
				if expectedCard.Code != drawnCards[i].Code {
					t.Errorf("Expected card code %s, but got %s", expectedCard.Code, drawnCards[i].Code)
				}
				if expectedCard.Value != drawnCards[i].Value {
					t.Errorf("Expected card value %s, but got %s", expectedCard.Value, drawnCards[i].Value)
				}
				if expectedCard.Suit != drawnCards[i].Suit {
					t.Errorf("Expected card suit %s, but got %s", expectedCard.Suit, drawnCards[i].Suit)
				}
			}
		})
	}
}
