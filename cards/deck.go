package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newdeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Aces", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	firstHand := d[:handSize]
	secondHand := d[handSize:]
	return firstHand, secondHand
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	strdeck := strings.Join([]string(d), ",")
	return strdeck
}

func (d deck) saveToFile(filename string) error {
	writestatus := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return writestatus
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// New source to generate a random seed each time, else random
	// sequence will always be the same.

	// Use timeNow to generate a different number each time.
	// UnixNano is int64 by default.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Position swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
