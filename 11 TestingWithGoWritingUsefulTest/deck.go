package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//joining a slice of strings
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Saving Data to the Hard Drive convert string(function toString)
//to write to byte
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

//Reading From the Hard Drive
func newDeckFromFile(filename string) deck {
	// bs is byte error
	bs, err := ioutil.ReadFile(filename)
	//Option #1 - log the error and return a call to newDeck()
	//Option #2 - log the error and entirely quite the program
	//jika error tidak kosong maka return error massage, err = nil berarti dapat baca file atau ada isinya

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//s string slice like deck
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
