package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck' which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "One", "Two"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "|")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	convertedStr := string(byteSlice)
	s := strings.Split(convertedStr, "|")
	return deck(s)
}

func (d deck) shuffle() {
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	for i := range d {
		rand.Seed(time.Now().UnixNano())
		swap := rand.Intn(len(d) - 1)

		d[i], d[swap] = d[swap], d[i]
	}
}
