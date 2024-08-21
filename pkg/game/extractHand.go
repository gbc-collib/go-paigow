package game

import (
	"sort"
)

func extractFlush(hand *[]*Card) []*Card {
	suitCount := getSuitCount(*hand)
	var flush []*Card
	for suit, count := range suitCount {
		if count >= 5 {
			for i, card := range *hand {
				if card.Suit == suit {
					flush = append(flush, card)
					*hand = append((*hand)[:i], (*hand)[i+1:]...)
				}
			}

		}
	}
	return flush
}

func sortHand(hand []*Card) []*Card {
	sort.Slice(hand, func(i, j int) bool {
		return hand[i].Value < hand[j].Value
	})
	return hand
}

func extractStraight(hand *[]*Card) []*Card {
	*hand = sortHand(*hand)
	var straight []*Card

	if isStraight(*hand) {
		count := 1
		straight = append(straight, (*hand)[0])

		for i := 1; i < len(*hand) && count < 5; i++ {
			if (*hand)[i].Value == (*hand)[i-1].Value+1 {
				straight = append(straight, (*hand)[i])
				count++
			} else if (*hand)[i].Value != (*hand)[i-1].Value {
				straight = []*Card{(*hand)[i]}
				count = 1
			}
		}

		if count == 5 {
			// Remove the cards in the straight from the original hand
			for _, card := range straight {
				for j := 0; j < len(*hand); j++ {
					if (*hand)[j].Value == card.Value && (*hand)[j].Suit == card.Suit {
						*hand = append((*hand)[:j], (*hand)[j+1:]...)
						break
					}
				}
			}
			return straight
		}
	}

	return nil
}

func extractFourOfKind(hand *[]*Card) []*Card {
	rankCount := getRankCount(*hand)
	var four []*Card
	for value, count := range rankCount {
		if count == 4 {
			for i, card := range *hand {
				if value == card.Value {
					four = append(four, card)
					*hand = append((*hand)[:i], (*hand)[i+1:]...)
				}
			}
			return four
		}
	}
	return nil

}

func extractThreeOfKind(hand *[]*Card) []*Card {
	rankCount := getRankCount(*hand)
	var three []*Card
	for value, count := range rankCount {
		if count == 3 {
			for i, card := range *hand {
				if value == card.Value {
					three = append(three, card)
					*hand = append((*hand)[:i], (*hand)[i+1:]...)
				}
			}
			return three
		}
	}
	return nil

}

func extractPair(hand *[]*Card) []*Card {
	if hand == nil || *hand == nil {
		return nil
	}
	rankCount := getRankCount(*hand)
	var pair []*Card
	for value, count := range rankCount {
		if count == 2 {
			for i := 0; i < len(*hand); {
				card := (*hand)[i]
				if value == card.Value {
					pair = append(pair, card)
					*hand = removeCard(*hand, card)
					// No increment here as the slice has been modified
				} else {
					i++
				}
			}
			return pair
		}
	}
	return nil
}

func extractFullHouse(hand *[]*Card) []*Card {
	//We don't want to extract until we know there is full house
	//A full house can be broken down into a three of kind and pair
	isThree := isThreeOfAKind(*hand)
	if isThree {
		pair := extractPair(hand)
		if pair != nil {
			//Now that we know there is afullhouse let's extract
			threes := extractThreeOfKind(hand)
			fullHouse := append(threes, pair...)
			return fullHouse
		}
	}
	return nil
}

func extractTwoPair(hand *[]*Card) []*Card {
	if isTwoPair(*hand) {
		twoPair := extractPair(hand)
		twoPair = append(twoPair, extractPair(hand)...)
		return twoPair
	}
	return nil
}

// TODO: There is probably a condition where this fails
// i.e. if we have 6 cards of a straight and the highest card is the card that makes it a flush
// Then this would return too soon and not find the straightFlush
func extractStraightFlush(hand *[]*Card) []*Card {
	if isStraightFlush(*hand) {
		straightFlush := extractStraight(hand)
		return straightFlush
	}
	return nil
}

func extractHighCard(hand *[]*Card) []*Card {
	// Sort the hand in descending order to get the highest card(s) first
	*hand = sortHand(*hand)
	var highCard []*Card
	// Find the highest card that is not nil
	for i := len(*hand) - 1; i >= 0; i-- {
		if (*hand)[i] != nil {
			highCard = append(highCard, (*hand)[i])
			break
		}
	}

	return highCard
}

func extractBestHand(hand *[]*Card) []*Card {
	if len(*hand) >= 4 {
		bestHand := extractStraightFlush(hand)
		if bestHand != nil {
			return bestHand
		}
		bestHand = extractFourOfKind(hand)
		if bestHand != nil {
			return bestHand
		}
		bestHand = extractFullHouse(hand)
		if bestHand != nil {
			return bestHand
		}
		bestHand = extractFlush(hand)
		if bestHand != nil {
			return bestHand
		}

		bestHand = extractStraight(hand)
		if bestHand != nil {
			return bestHand
		}
	}
	if len(*hand) >= 3 {
		bestHand := extractThreeOfKind(hand)
		if bestHand != nil {
			return bestHand
		}
	}
	if len(*hand) >= 4 {
		bestHand := extractTwoPair(hand)
		if bestHand != nil {
			return bestHand
		}
	}
	bestHand := extractPair(hand)
	if bestHand != nil {
		return bestHand
	}
	bestHand = extractHighCard(hand)
	return bestHand
}
