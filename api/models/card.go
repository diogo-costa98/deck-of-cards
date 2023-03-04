package models

import (
	"strconv"
)

// Card represents a playing card.
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// Creates a list of playing cards based on their codes, if they exist
func CreateCards(cardCodes []string) []Card {
	cards := []Card{}

	if len(cardCodes) == 0 {
		for _, suit := range []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"} {
			cards = append(cards, Card{
				Value: "ACE",
				Suit:  suit,
				Code:  "A" + suit[0:1],
			})
			for i := 2; i <= 10; i++ {
				cards = append(cards, Card{
					Value: strconv.Itoa(i),
					Suit:  suit,
					Code:  strconv.Itoa(i) + suit[0:1],
				})
			}
			cards = append(cards, Card{
				Value: "JACK",
				Suit:  suit,
				Code:  "J" + suit[0:1],
			})
			cards = append(cards, Card{
				Value: "QUEEN",
				Suit:  suit,
				Code:  "Q" + suit[0:1],
			})
			cards = append(cards, Card{
				Value: "KING",
				Suit:  suit,
				Code:  "K" + suit[0:1],
			})
		}
		return cards
	}

	for _, suit := range []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"} {
		for _, value := range []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"} {
			for _, code := range cardCodes {
				if code == value[0:1]+suit[0:1] {
					cards = append(cards, Card{
						Value: value,
						Suit:  suit,
						Code:  code,
					})
					break
				}
			}
		}
	}

	return cards
}
