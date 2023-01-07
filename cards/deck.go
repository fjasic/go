package main

import "fmt"

// create new type of deck which is a slice of strings
type deck []string

// reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func new_deck() deck {
	cards := deck{}
	cards_suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cards_values := []string{"Ace", "Two", "Tree", "Four"}
	for _, suit := range cards_suits {
		for _, value := range cards_values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, hand_size int) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}
