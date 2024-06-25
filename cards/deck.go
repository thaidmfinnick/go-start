package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			mergedString := fmt.Sprint(suit, " ", value)
			cards = append(cards, mergedString)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	first := d[:handSize]
	second := d[handSize:]
	return first, second
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
