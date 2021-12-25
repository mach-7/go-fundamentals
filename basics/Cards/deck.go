package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Convert the deck of cards to a string representation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save the deck of cards to a text file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create a new deck of cards from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 log the error and exit
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(bs), ",")
	return deck(stringSlice)
}

// Shuffle the deck of cards

func (d deck) shuffle() {
	// generate some random numbers and swap the cards in the deck using that
	randomNumberSeed := time.Now().UnixNano()
	source := rand.NewSource(randomNumberSeed)
	r := rand.New(source)

	for i := range d {
		// generate a random number between 0 and (length of Deck - 1)
		newPosition := r.Intn(len(d) - 1)
		// fancy syntax to swap two positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
