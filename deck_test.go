package deck

import (
	"os"
	"reflect"
	"testing"
)

func TestGetCardSuits(t *testing.T) {

	if len(getCardSuits()) != 4 {
		t.Errorf("Expected 4 card suits and got %v", len(getCardSuits()))
	}
}

func TestGetCardValues(t *testing.T) {

	if len(getCardValues()) != 13 {
		t.Errorf("Expected 13 card values and got %v", len(getCardValues()))
	}
}

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 card in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	shuffledDeck := newDeck()
	shuffledDeck.shuffle()

	if reflect.DeepEqual(d, shuffledDeck) {
		t.Errorf("The deck wasn't shuffled")
	}

}

func TestDeal(t *testing.T) {
	d := newDeck()

	_, d, err := deal(d, 53)

	if err == nil {
		t.Errorf("Expected error when hand is greather then deck")
	}

	hand, d, _ := deal(d, 5)

	if len(d) == 52 {
		t.Errorf("The length of the deck shout be less then 52")
	}

	if len(hand) != 5 {
		t.Errorf("The length of the hand should be 5")
	}

}
