package main

func main() {
	cards := newDeck()
	cards = append(cards, "six of Spades")
	cards.print()
}
