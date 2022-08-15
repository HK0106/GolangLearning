package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, suit+" of "+value)
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

func toString(d deck) string {
	return strings.Join([]string(d), ",")
}

func toByteSlice(d deck) []byte {
	return []byte(toString(d))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, toByteSlice(d), 0666)
}

func readFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), ","))
}

func (d deck) shuffle() {
	for i := range d {
		newPositon := rand.Intn(len(d) - 1)
		d[i], d[newPositon] = d[newPositon], d[i]
	}

}
