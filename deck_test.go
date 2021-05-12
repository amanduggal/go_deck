package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected different first card. Got %v", d[0])
	}

	if d[len(d)-1] != "Two of Clubs" {
		t.Errorf("Expected different last card. Got %v", d[len(d)-1])
	}
}

func TestFilesIO(t *testing.T) {
	fileNameToSaveAndRead := "_decktesting"
	os.Remove(fileNameToSaveAndRead)

	d := newDeck()

	hand, _ := deal(d, 5)
	if len(hand) != 5 {
		t.Errorf("Expected length of hand dealt to be 5. Got %v", len(hand))
	}

	hand.saveToFile(fileNameToSaveAndRead)

	readDeck := readFile(fileNameToSaveAndRead)
	if len(readDeck) != 5 {
		t.Errorf("Expected length of hand dealt to be 5. Got %v", len(readDeck))
	}

	if hand[0] != readDeck[0] {
		t.Errorf("First elements don't match ")
	}
}
