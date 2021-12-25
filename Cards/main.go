package main

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
	cards.print()
}
