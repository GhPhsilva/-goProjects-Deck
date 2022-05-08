package deck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func getCardSuits() [4]string {
	return [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
}

func getCardValues() [13]string {
	return [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Queen", "Joker", "King"}
}

func newDeck() deck {
	cards := deck{}
	suits := getCardSuits()
	values := getCardValues()

	for _, value := range values {
		for _, suit := range suits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck, error) {
	if handSize > len(d) {
		return nil, d, errors.New("the handsize is greather then the rest of cards in deck")
	}

	return d[:handSize], d[handSize:], nil
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
