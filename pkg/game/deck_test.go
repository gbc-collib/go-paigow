package game

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	//Paigow is 52 card deck + Joker
	if len(deck.Cards) != 53 {
		t.Errorf("expected 53 cards, got %d", len(deck.Cards))
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	originalOrder := make([]*Card, len(deck.Cards))
	copy(originalOrder, deck.Cards)
	deck.Shuffle()
	shuffledOrder := deck.Cards
	sameOrder := true
	for i, card := range shuffledOrder {
		if card != originalOrder[i] {
			sameOrder = false
			break
		}
	}
	if sameOrder {
		t.Error("expected shuffled order to be different from original order")
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	want := 7
	got := deck.Draw(7)
	if len(got) != want {
		t.Errorf("Got %v, Wanted %v", got, want)
	}
}
