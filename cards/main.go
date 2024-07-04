package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, _ := deal(cards, 5)
	// hand.saveToFile("my_hand")
	test := newDeckFromFile("my_hand")
	fmt.Println(test)
}
