package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	len := len(d)

	if len != 16 {
		t.Errorf("Error, expected a value of 16, but got %v", len)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Error, expected a value of Ace Of Spades, but got %v", d[0])
	}

	if d[len-1] != "Four of Clubs" {
		t.Errorf("Error, expected a value of Four Of Clubs, but got %v", d[len-1])
	}
}

func TestWriteAndReadFromFile(t *testing.T) {

	os.Remove("_deck_testing")
	d := newDeck()

	d.saveToFile("_deck_testing")

	newDeck := newDeckFromFile("_deck_testing")

	len := len(newDeck)

	if len != 16 {
		t.Errorf("Error, expected size of deck to be 16, but got %v", len)
	}

	os.Remove("_deck_testing")

}
