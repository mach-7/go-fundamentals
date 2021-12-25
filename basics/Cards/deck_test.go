package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Testing if the deck has 52 cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %d", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// create temporary file _decktesting.txt
	filename := "_decktesting.txt"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected a deck with length 52 but got %d", len(loadedDeck))
	}

}
