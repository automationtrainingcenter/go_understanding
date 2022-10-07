package main

import "fmt"

func main() {
	// card := newCard()
	// fmt.Println(card)
	// getCards()
	// use the newly created deck type
	cards := newDeck()
	cards = append(cards, newCard())
	cards.print()

}

func newCard() string {
	return "King of Diamonds"
}

func getCards() {
	cards := []string {"Ace of Diamonds", "Two of Spades"} // create a Slice
	// Slice is a dynamic array
	// add a new card to Slice
	cards = append(cards, newCard())

	// iterate over the cards Slice using for loop
	for i, card := range cards {
		fmt.Println(i, card)
	}

}