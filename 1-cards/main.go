package main

func main() {
	cards := newDeck()
	// // cards.saveToFile("myCards.txt")
	// cards := newDeckFromFile("myCards.txt")
	cards.shuffle()
	cards.print()
}
