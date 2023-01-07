package main

import "fmt"

func main() {
	cards := new_deck()
	hand, remaining_deck := deal(cards, 5)

	fmt.Printf("----- Hand -----\n")
	hand.print()
	fmt.Printf("----- Remaining deck -----\n")
	remaining_deck.print()

}
