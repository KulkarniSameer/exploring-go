package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T)  {
	d := newDeck()

	dLen := len(d)
	expectedLength := 52

	if dLen != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, dLen)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T)  {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(d) != len(loadedDeck) {
		t.Errorf("decks do not match")
	}

	os.Remove("_decktesting")
}