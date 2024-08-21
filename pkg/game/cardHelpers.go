package game

func getRankCount(hand []*Card) map[Value]int {
	rankCount := make(map[Value]int)
	for _, card := range hand {
		rankCount[card.Value]++
	}
	return rankCount
}

func getSuitCount(hand []*Card) map[Suit]int {
	suitCount := make(map[Suit]int)
	for _, card := range hand {
		suitCount[card.Suit]++
	}
	return suitCount
}

func getRanks(hand []*Card) []Value {
	ranks := make([]Value, len(hand))
	for index, card := range hand {
		ranks[index] = card.Value
	}
	return ranks
}

func getMatchingSuits(hand []*Card) []*Card {
	count := getSuitCount(hand)
	highestCount := 0
	var highestSuit Suit
	var matching []*Card
	for suit, value := range count {
		if value > highestCount {
			highestSuit = suit
		}
	}
	for _, card := range hand {
		if card.Suit == highestSuit {
			matching = append(matching, card)
		}
	}
	return matching

}

func convertToFiveCardArray(cards []*Card) [5]*Card {
	var arr [5]*Card
	for i := 0; i < 5 && i < len(cards); i++ {
		arr[i] = cards[i]
	}
	return arr
}

func convertToTwoCardArray(cards []*Card) [2]*Card {
	var arr [2]*Card
	for i := 0; i < 2 && i < len(cards); i++ {
		arr[i] = cards[i]
	}
	return arr
}

// TODO: refactor to use this everywhere
func removeCard(cards []*Card, cardToRemove *Card) []*Card {
	for i, card := range cards {
		if card == cardToRemove {
			return append(cards[:i], cards[i+1:]...)
		}
	}
	return cards
}
