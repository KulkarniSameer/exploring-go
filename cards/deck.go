package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which a slice of string
type deck []string

func newDeck() deck {
	cards := deck{};
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " Of " + suit)
		}
	}

	return cards;
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	var d deck;
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	d = deck([]string(strings.Split(string(byteSlice), ",")))
	return d
}

func (d deck) print()  {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) deal(handSize int) deck {
	return d[:handSize]
}

func (d deck) shuffle() deck {
	// create a custom source to avoid same seed value for
	// random number generator.
	source := rand.NewSource(time.Now().UnixNano());

	// shuffle the cards in the deck
	// iterate over all the cards, for every entry generate a number
	// shuffle the cards
	for i := range d {
		newPos := rand.New(source).Intn((len(d) - 1))
		d[i], d[newPos] = d[newPos], d[i]
	}
	return d
}

func (d deck) saveToFile(filename string) error {
	byteSlice := []byte(d.toString())

	return ioutil.WriteFile(filename, byteSlice, 0666)
}