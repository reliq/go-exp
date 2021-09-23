package main

import (
	"os"
	"testing"
)

var expectedLen int

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen = 52
	expectedFirstEl := "Ace of Spades"
	expectedLastEl := "Jack of Diamonds"

	if len(d) != expectedLen {
		t.Errorf("Expected deck length to be %d, but got %v", expectedLen, len(d))
	}

	if d[0] != expectedFirstEl {
		t.Errorf("Expected first card of %v. Got %s", expectedFirstEl, d[0])
	}

	if d[len(d)-1] != expectedLastEl {
		t.Errorf("Expected last card of %v. Got %s", expectedLastEl, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	tempFile := "_decktesting"
	expectedLen = 52

	os.Remove(tempFile)

	deck := newDeck()
	deck.saveToFile(tempFile)

	loadedDeck := newDeckFromFile(tempFile)

	if len(loadedDeck) != expectedLen {
		t.Errorf("Expected %d cards in deck, got %v", expectedLen, len(loadedDeck))
	}

	os.Remove(tempFile)
}
