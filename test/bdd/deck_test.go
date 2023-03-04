package bdd

import (
	"github.com/diogo-costa98/deck-of-cards/api/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deck", func() {
	var (
		deck         *models.Deck
		ok           bool
		drawnCards   []models.Card
		shuffledDeck *models.Deck
	)

	BeforeEach(func() {
		deck = models.CreateDeck(nil, false)
	})

	Describe("Creating a deck of cards", func() {
		It("should have 52 cards", func() {
			Expect(len(deck.Cards)).To(Equal(52))
		})

		It("should be unshuffled", func() {
			Expect(deck.Shuffled).To(BeFalse())
		})

		It("should have a remaining count of 52", func() {
			Expect(deck.Remaining).To(Equal(52))
		})
	})

	Describe("Drawing cards from the deck", func() {
		BeforeEach(func() {
			drawnCards, ok = models.DrawDeck(5, deck)
		})

		It("should receive 5 cards", func() {
			Expect(len(drawnCards)).To(Equal(5))
		})
	})

	Describe("Drawing more cards than exist in the deck", func() {
		BeforeEach(func() {
			drawnCards, ok = models.DrawDeck(53, deck)
		})

		It("should return an error", func() {
			Expect(ok).To(BeFalse())
		})
	})

	Describe("Create shuffled deck", func() {
		BeforeEach(func() {
			shuffledDeck = models.CreateDeck(nil, true)
		})

		It("should be shuffled", func() {
			Expect(shuffledDeck.Shuffled).To(BeTrue())
		})

		It("should have a different order of cards", func() {
			Expect(shuffledDeck.Cards).ToNot(Equal(deck.Cards))
		})
	})
})
