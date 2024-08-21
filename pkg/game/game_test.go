package game

import (
	"reflect"
	"testing"
)

func dereferenceCards(cards []*Card) []Card {
	result := make([]Card, len(cards))
	for i, card := range cards {
		if card != nil {
			result[i] = *card
		}
	}
	return result
}
func dereferenceHands(hands Hands) struct {
	HighHand []Card
	LowHand  []Card
} {
	return struct {
		HighHand []Card
		LowHand  []Card
	}{
		HighHand: dereferenceCards(hands.HighHand[:]),
		LowHand:  dereferenceCards(hands.LowHand[:]),
	}
}

func TestAddPlayer(t *testing.T) {
	t.Run("Add Too Many Players", func(t *testing.T) {
		game := NewGame()
		err := game.AddPlayers(7)
		assertError(t, err, ErrTooManyPlayers)
	})

	t.Run("Add Players", func(t *testing.T) {
		game := NewGame()
		err := game.AddPlayers(2)
		if err != nil {
			t.Errorf("Got error: %#v when none expected", err)
		}
		want := 2
		got := game.CountPlayers()
		if got != want {
			t.Errorf("Got %v, Wanted %v", got, want)
		}

	})

}

func TestDeal(t *testing.T) {
	game := NewGame()
	game.AddPlayers(1)
	game.Deal()
	got := game.Players[0].Hand
	for _, card := range got {
		if card == nil {
			t.Errorf("Wanted not nil for Card, Got %v", card)
		}
	}
	if len(game.Deck.Cards) > 0 {
		t.Errorf("Deck Should be empty, %v cards Left", len(game.Deck.Cards))
	}
}



func TestEvalauteHands(t *testing.T) {
	tests := []struct {
		hand     []*Card
		expected HandRank
	}{
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 4, Suit: Diamond},
				{Value: 5, Suit: Heart},
				{Value: 10, Suit: Heart},
				{Value: 2, Suit: Spade},
				{Value: 8, Suit: Club},
				{Value: 14, Suit: Club},
			},
			expected: OnePair,
		},
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 3, Suit: Spade},
				{Value: 4, Suit: Diamond},
				{Value: 5, Suit: Heart},
				{Value: 6, Suit: Heart},
			},
			expected: Straight,
		},
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 3, Suit: Heart},
				{Value: 4, Suit: Heart},
				{Value: 5, Suit: Heart},
				{Value: 6, Suit: Heart},
				{Value: 9, Suit: Spade},
				{Value: 10, Suit: Spade},
			},
			expected: StraightFlush,
		},

		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 2, Suit: Spade},
				{Value: 2, Suit: Diamond},
				{Value: 5, Suit: Heart},
				{Value: 6, Suit: Heart},
			},
			expected: ThreeOfAKind,
		},
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 2, Suit: Spade},
				{Value: 2, Suit: Diamond},
				{Value: 3, Suit: Heart},
				{Value: 3, Suit: Heart},
			},
			expected: FullHouse,
		},
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 2, Suit: Spade},
				{Value: 2, Suit: Diamond},
				{Value: 3, Suit: Heart},
				{Value: 3, Suit: Heart},
				{Value: 4, Suit: Club},
				{Value: 8, Suit: Club},
			},
			expected: FullHouse,
		},
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 2, Suit: Spade},
				{Value: 5, Suit: Diamond},
				{Value: 3, Suit: Heart},
				{Value: 3, Suit: Heart},
				{Value: 8, Suit: Club},
				{Value: 8, Suit: Club},
			},
			expected: TwoPair,
		},
		{
			hand: []*Card{
				{Value: 10, Suit: Heart},
				{Value: 2, Suit: Spade},
				{Value: 5, Suit: Diamond},
				{Value: 8, Suit: Heart},
				{Value: 3, Suit: Heart},
				{Value: 4, Suit: Club},
				{Value: 9, Suit: Club},
			},
			expected: HighCard,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := EvaluateHighHand(tt.hand)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

func TestFindBestHands(t *testing.T) {
	tests := []struct {
		hand     []*Card
		expected Hands
	}{
		{
			hand: []*Card{
				{Value: 2, Suit: Heart},
				{Value: 3, Suit: Heart},
				{Value: 4, Suit: Heart},
				{Value: 5, Suit: Heart},
				{Value: 6, Suit: Heart},
				{Value: 7, Suit: Diamond},
				{Value: 8, Suit: Club},
			},
			expected: Hands{
				LowHand: [2]*Card{
					{Value: 8, Suit: Club},
					{Value: 7, Suit: Diamond},
				},
				HighHand: [5]*Card{
					{Value: 2, Suit: Heart},
					{Value: 3, Suit: Heart},
					{Value: 4, Suit: Heart},
					{Value: 5, Suit: Heart},
					{Value: 6, Suit: Heart},
				},
			},
		},
		{
			hand: []*Card{
				{Value: 10, Suit: Spade},
				{Value: Jack, Suit: Spade},
				{Value: Queen, Suit: Spade},
				{Value: King, Suit: Spade},
				{Value: Ace, Suit: Spade},
				{Value: 2, Suit: Diamond},
				{Value: 3, Suit: Club},
			},
			expected: Hands{
				LowHand: [2]*Card{
					{Value: 3, Suit: Club},
					{Value: 2, Suit: Diamond},
				},
				HighHand: [5]*Card{
					{Value: 10, Suit: Spade},
					{Value: Jack, Suit: Spade},
					{Value: Queen, Suit: Spade},
					{Value: King, Suit: Spade},
					{Value: Ace, Suit: Spade},
				},
			},
		},
		{
			hand: []*Card{
				{Value: 2, Suit: Club},
				{Value: 2, Suit: Diamond},
				{Value: 3, Suit: Heart},
				{Value: 3, Suit: Spade},
				{Value: 4, Suit: Club},
				{Value: 4, Suit: Diamond},
				{Value: 9, Suit: Heart},
			},
			expected: Hands{
				LowHand: [2]*Card{
					{Value: 4, Suit: Club},
					{Value: 4, Suit: Diamond},
				},
				HighHand: [5]*Card{
					{Value: 2, Suit: Club},
					{Value: 2, Suit: Diamond},
					{Value: 3, Suit: Heart},
					{Value: 3, Suit: Spade},
					{Value: 9, Suit: Heart},
				},
			},
		},
	}

	for i, tt := range tests {
		if i == 2 {
			t.Run("", func(t *testing.T) {
				result := FindBestHands(&tt.hand)
				if !reflect.DeepEqual(dereferenceHands(result), dereferenceHands(tt.expected)) {
					t.Errorf("expected %+v, got %+v", dereferenceHands(tt.expected), dereferenceHands(result))
				}
			})
		}
	}
}
