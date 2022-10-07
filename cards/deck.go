package main

import "fmt"

// create a new type of deck which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	// creates a new deck and returns it
	cards := deck {}

	cardSuites := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of "+ suite)
		}
	}

	return cards
}
