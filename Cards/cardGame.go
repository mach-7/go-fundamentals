package main

import "fmt"

/*
A basic game of Cards
- Create a new Deck
- Print
- Shuffle
- Deal some cards
- Save a list of cards to file
- Load a list of cards from file
*/

func main() {
	var card string = newCard()

	cardSlice := []string{"9 of Diamonds", "Jack of Clubs"}

	fmt.Println(card)

	for i, card := range cardSlice {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}

func newDeck() {

}

func printDeck() {

}

func shuffle() {

}

func dealCards() {

}

func saveDeck() {

}

func loadDeck() {

}
