package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createNewDeck() deck {
	cards := deck{}

	suitCards := []string{"Heart", "Diamond", "Spade", "Club"}
	valueCards := []string{"One", "Two", "Three", "Four"}

	for _, suit := range suitCards {
		for _, value := range valueCards {
			cards = append(cards, value+" "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for idx, card := range d {
		fmt.Println(string(idx) + ". " + card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func deal(d deck, hand int) (lefthand deck, righthand deck) {
	return d[0:hand], d[hand:]
}

func saveToFile(d deck) {
	ioutil.WriteFile("deck", []byte(strings.Join([]string(d), ",")), 0777)
}

func newDeckFromFile(filename string) (newDeck deck) {
	content, error := ioutil.ReadFile(filename)
	if error != nil {
		log.Println("Cannot read file.")
		os.Exit(1)
	}
	return deck(strings.Split(string(content), ","))
}
