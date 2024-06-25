package main

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)
	hand.saveToFile("my_hand")
}
