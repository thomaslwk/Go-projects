package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newdeck()
	// Check total deck size
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	// Check first card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
	}
	// Check last card
	if d[len(d)-1] != "Queen of hearts" {
		t.Errorf("Expected last card to be Queen of hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove existing file if present
	os.Remove("_decktesting")

	// Create new deck
	deck := newdeck()
	deck.saveToFile("_decktesting")

	// load new deck file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	// Remove testing deck
	os.Remove("_decktesting")
}
