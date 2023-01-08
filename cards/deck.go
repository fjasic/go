package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of deck which is a slice of strings
type deck []string

// reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range d {
		np := rnd.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
	d.print()
}
