package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Length is not as expected, got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("First card is not Ace of Spades as expected, got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Last card not Four of Clubs as expected, got %v", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_deckstring")
	d := newDeck()
	d.saveToFile("_decktesting")
	nd := newDeckFromFile("_decktesting")
	if len(nd) != len(d) {
		t.Errorf("Save to file does not work, got %v expected 16 cards!", len(nd))
	}
	os.Remove("_deckstring")
}
