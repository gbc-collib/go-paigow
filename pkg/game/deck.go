package game

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []*Card
}

func NewDeck() *Deck {
	var cards []*Card
	for suit := Club; suit <= Spade; suit++ {
		for value := Two; value <= Ace; value++ {
			cards = append(cards, &Card{Suit: Suit(suit), Value: value})
		}
	}
	cards = append(cards, &Card{Suit: Suit(0), Value: 0})
	return &Deck{Cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw(n int) []*Card {
	drawn := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return drawn
}
