package game

import (
	"testing"
)

func TestEvaluateLowHand(t *testing.T) {
	tests := []struct {
		hand     [2]*Card
		expected HandRank
	}{
		{
			hand: [2]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Two},
			},
			expected: OnePair,
		},
		{
			hand: [2]*Card{
				{Suit: Spade, Value: Three},
				{Suit: Heart, Value: Two},
			},
			expected: HighCard,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := evaluateLowHand(tt.hand)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEvaluateHighHand(t *testing.T) {
	tests := []struct {
		hand     [5]*Card
		expected HandRank
	}{
		// Test High Card
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Four},
				{Suit: Diamond, Value: Six},
				{Suit: Club, Value: Eight},
				{Suit: Spade, Value: Ten},
			},
			expected: HighCard,
		},
		// Test One Pair
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Two},
				{Suit: Diamond, Value: Six},
				{Suit: Club, Value: Eight},
				{Suit: Spade, Value: Ten},
			},
			expected: OnePair,
		},
		// Test Two Pair
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Two},
				{Suit: Diamond, Value: Six},
				{Suit: Club, Value: Six},
				{Suit: Spade, Value: Ten},
			},
			expected: TwoPair,
		},
		// Test Three of a Kind
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Two},
				{Suit: Diamond, Value: Two},
				{Suit: Club, Value: Eight},
				{Suit: Spade, Value: Ten},
			},
			expected: ThreeOfAKind,
		},
		// Test Straight
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Three},
				{Suit: Diamond, Value: Four},
				{Suit: Club, Value: Five},
				{Suit: Spade, Value: Six},
			},
			expected: Straight,
		},
		// Test Flush
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Spade, Value: Four},
				{Suit: Spade, Value: Six},
				{Suit: Spade, Value: Eight},
				{Suit: Spade, Value: Ten},
			},
			expected: Flush,
		},
		// Test Full House
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Two},
				{Suit: Diamond, Value: Two},
				{Suit: Club, Value: Four},
				{Suit: Spade, Value: Four},
			},
			expected: FullHouse,
		},
		// Test Four of a Kind
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Heart, Value: Two},
				{Suit: Diamond, Value: Two},
				{Suit: Club, Value: Two},
				{Suit: Spade, Value: Ten},
			},
			expected: FourOfAKind,
		},
		// Test Straight Flush
		{
			hand: [5]*Card{
				{Suit: Spade, Value: Two},
				{Suit: Spade, Value: Three},
				{Suit: Spade, Value: Four},
				{Suit: Spade, Value: Five},
				{Suit: Spade, Value: Six},
			},
			expected: StraightFlush,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := EvaluateHighHand(tt.hand[:])
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
