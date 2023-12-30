package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Heart")

	for index, card := range cards {
		fmt.Println(index, card)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i+1, cards[i])
	}
}

func newCard() string {
	return "Five of Dimonds"
}
