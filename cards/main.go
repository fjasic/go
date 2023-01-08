package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards")
	deck := newDeckFromFile("cards")
	deck.print()
	deck.shuffle()
}
