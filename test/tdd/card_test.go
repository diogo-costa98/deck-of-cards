package tdd

import (
	"testing"

	"github.com/diogo-costa98/deck-of-cards/api/models"
)

func TestCreateCards(t *testing.T) {
	testCases := []struct {
		name          string
		cardCodes     []string
		expectedCards []models.Card
	}{
		{
			name:      "test with nil card codes",
			cardCodes: nil,
			expectedCards: []models.Card{
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
				{Value: "2", Suit: "SPADES", Code: "2S"},
				{Value: "3", Suit: "SPADES", Code: "3S"},
				{Value: "4", Suit: "SPADES", Code: "4S"},
				{Value: "5", Suit: "SPADES", Code: "5S"},
				{Value: "6", Suit: "SPADES", Code: "6S"},
				{Value: "7", Suit: "SPADES", Code: "7S"},
				{Value: "8", Suit: "SPADES", Code: "8S"},
				{Value: "9", Suit: "SPADES", Code: "9S"},
				{Value: "10", Suit: "SPADES", Code: "10S"},
				{Value: "JACK", Suit: "SPADES", Code: "JS"},
				{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
				{Value: "KING", Suit: "SPADES", Code: "KS"},
				{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
				{Value: "2", Suit: "DIAMONDS", Code: "2D"},
				{Value: "3", Suit: "DIAMONDS", Code: "3D"},
				{Value: "4", Suit: "DIAMONDS", Code: "4D"},
				{Value: "5", Suit: "DIAMONDS", Code: "5D"},
				{Value: "6", Suit: "DIAMONDS", Code: "6D"},
				{Value: "7", Suit: "DIAMONDS", Code: "7D"},
				{Value: "8", Suit: "DIAMONDS", Code: "8D"},
				{Value: "9", Suit: "DIAMONDS", Code: "9D"},
				{Value: "10", Suit: "DIAMONDS", Code: "10D"},
				{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
				{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
				{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
				{Value: "ACE", Suit: "CLUBS", Code: "AC"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "3", Suit: "CLUBS", Code: "3C"},
				{Value: "4", Suit: "CLUBS", Code: "4C"},
				{Value: "5", Suit: "CLUBS", Code: "5C"},
				{Value: "6", Suit: "CLUBS", Code: "6C"},
				{Value: "7", Suit: "CLUBS", Code: "7C"},
				{Value: "8", Suit: "CLUBS", Code: "8C"},
				{Value: "9", Suit: "CLUBS", Code: "9C"},
				{Value: "10", Suit: "CLUBS", Code: "10C"},
				{Value: "JACK", Suit: "CLUBS", Code: "JC"},
				{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
				{Value: "KING", Suit: "CLUBS", Code: "KC"},
				{Value: "ACE", Suit: "HEARTS", Code: "AH"},
				{Value: "2", Suit: "HEARTS", Code: "2H"},
				{Value: "3", Suit: "HEARTS", Code: "3H"},
				{Value: "4", Suit: "HEARTS", Code: "4H"},
				{Value: "5", Suit: "HEARTS", Code: "5H"},
				{Value: "6", Suit: "HEARTS", Code: "6H"},
				{Value: "7", Suit: "HEARTS", Code: "7H"},
				{Value: "8", Suit: "HEARTS", Code: "8H"},
				{Value: "9", Suit: "HEARTS", Code: "9H"},
				{Value: "10", Suit: "HEARTS", Code: "10H"},
				{Value: "JACK", Suit: "HEARTS", Code: "JH"},
				{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
			},
		},
		{
			name:      "test with partial card codes",
			cardCodes: []string{"AS", "KD", "AC", "2C", "KH"},
			expectedCards: []models.Card{
				{Value: "ACE", Suit: "SPADES", Code: "AS"},
				{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
				{Value: "ACE", Suit: "CLUBS", Code: "AC"},
				{Value: "2", Suit: "CLUBS", Code: "2C"},
				{Value: "KING", Suit: "HEARTS", Code: "KH"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cards := models.CreateCards(tc.cardCodes)
			if len(cards) != len(tc.expectedCards) {
				t.Fatalf("expected %d cards, but got %d", len(tc.expectedCards), len(cards))
			}
			for i, expectedCard := range tc.expectedCards {
				if expectedCard.Code != cards[i].Code {
					t.Errorf("expected card code %s, but got %s", expectedCard.Code, cards[i].Code)
				}
				if expectedCard.Value != cards[i].Value {
					t.Errorf("expected card value %s, but got %s", expectedCard.Value, cards[i].Value)
				}
				if expectedCard.Suit != cards[i].Suit {
					t.Errorf("expected card suit %s, but got %s", expectedCard.Suit, cards[i].Suit)
				}
			}
		})
	}
}
