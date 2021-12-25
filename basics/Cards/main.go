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
	cards := newDeck()
	hand, _ := deal(cards, 5)
	hand.print()
	// Print the hand to the screen
	fmt.Println(hand.toString())
	// Save the deck to a file
	err := cards.saveToFile("my_cards.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Saved To File ")
	}

	cards2 := newDeckFromFile("my_cards.txt")
	fmt.Println(cards2.toString())
}
