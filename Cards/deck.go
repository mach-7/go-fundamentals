package main

import "fmt"

// Create a custom type of 'deck' which is nothing but a slice [] of string
type deck []string

// New deck creation function
func newDeck() deck {
	cardSuites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// define a receiver function for the deck data type, that prints the cards in the deck
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// deal n number of cards
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
