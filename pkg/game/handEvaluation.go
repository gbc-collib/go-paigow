package game

import (
	"slices"
)

func EvaluateHighHand(hand []*Card) HandRank {
	if len(hand) >= 5 {
		if isStraightFlush(hand) {
			return StraightFlush
		}
	}
	if len(hand) >= 4 {
		if isFourOfAKind(hand) {
			return FourOfAKind
		}
	}
	if len(hand) >= 5 {
		if isFullHouse(hand) {
			return FullHouse
		}
		if isFlush(hand) {
			return Flush
		}
		if isStraight(hand) {
			return Straight
		}
	}
	if len(hand) >= 3 {
		if isThreeOfAKind(hand) {
			return ThreeOfAKind
		}
	}
	if len(hand) >= 4 {
		if isTwoPair(hand) {
			return TwoPair
		}
	}
	if len(hand) >= 2 {
		if isOnePair(hand) {
			return OnePair
		}
	}
	return HighCard
}

func isStraightFlush(hand []*Card) bool {
	return isFlush(hand) && isStraight(hand)
}

func isFullHouse(hand []*Card) bool {
	return isOnePair(hand) && isThreeOfAKind(hand)
}

func isFlush(hand []*Card) bool {
	suitCount := getSuitCount(hand)
	for _, count := range suitCount {
		if count == 5 {
			return true
		}
	}
	return false
}

func isThreeOfAKind(hand []*Card) bool {
	rankCount := getRankCount(hand)
	for _, count := range rankCount {
		if count == 3 {
			return true
		}
	}
	return false
}

func isFourOfAKind(hand []*Card) bool {
	rankCount := getRankCount(hand)
	for _, count := range rankCount {
		if count == 4 {
			return true
		}
	}
	return false
}

func isStraight(hand []*Card) bool {
	ranks := getRanks(hand)
	slices.Sort(ranks)

	// Check if Ace is present and handle it as both low (1) and high (14)
	hasAce := false
	for _, rank := range ranks {
		if rank == 14 {
			hasAce = true
			break
		}
	}

	// Include Ace as 1 for low straight checking
	if hasAce {
		ranks = append(ranks, 1)
	}

	count := 1
	maxCount := 1

	for i := 0; i < len(ranks)-1; i++ {
		if ranks[i+1] == ranks[i]+1 {
			count++
		} else if ranks[i+1] != ranks[i] { // Avoid counting duplicates
			count = 1
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount >= 5
}

func isTwoPair(hand []*Card) bool {
	rankCount := getRankCount(hand)
	pairCount := 0
	for _, count := range rankCount {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount >= 2
}

func isOnePair(hand []*Card) bool {
	rankCount := getRankCount(hand)
	for _, count := range rankCount {
		if count == 2 {
			return true
		}
	}
	return false
}

func evaluateLowHand(hand [2]*Card) HandRank {
	if hand[0].Value == hand[1].Value {
		return OnePair
	}
	return HighCard
}
