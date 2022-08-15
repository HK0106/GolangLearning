package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()

	//hand, deck := deal(cards, 5)
	//hand.print()
	//deck.print()

	//cards.saveToFile("my_Deck_1")

	//cards := readFromFile("my_Deck_11")
	//cards.print()

}
