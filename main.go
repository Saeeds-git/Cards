package main

func main() {
	cards := newDeckFromFile("myCards")
	cards.print()

	cards1 := newDeck()
	cards1.shuffle()
	cards1.print()
}
