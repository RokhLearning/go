package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("[Error] : lenth of Deck is not 52, it's %v", len(d))
	}
	if d[0] != "Ace of Diamond" {
		t.Errorf("[Error] : Expected the first card to be Ace of Diamonds, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("[Error] : Expected the last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("[Error] : lenth of Deck is not 52, it's %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
