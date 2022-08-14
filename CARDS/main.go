package main

func main() {
	cards := deck{"ASce of diamonds", newCard()}
	cards = append(cards, "Six of spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
