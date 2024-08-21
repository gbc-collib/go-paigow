package game

type Game struct {
	Deck    *Deck
	Dealer  *Player
	Players [6]*Player
}

type GameErr string

func (e GameErr) Error() string {
	return string(e)
}

const (
	ErrTooManyPlayers = GameErr("Could not Add Player, max Players is 6")
)

func NewGame() *Game {
	deck := NewDeck()
	deck.Shuffle()
	return &Game{Deck: deck, Dealer: newPlayer()}
}

func (g *Game) NewRound() {
	//1. Get New Deck and Shuffle
	deck := NewDeck()
	deck.Shuffle()
	g.Deck = deck
	//2. Deal Cards out
	g.Deal()
	//3. Have Dealer Solve Hand
	g.PlayDealerHand()
	//4 Allow Players to Play
}

func (g *Game) CountPlayers() int {
	count := 0
	for _, player := range g.Players {
		if player != nil {
			count++
		}
	}
	return count
}

func (g *Game) AddPlayers(numberOfPlayers int) error {
	//If Total Players greater than 6 Players
	if numberOfPlayers+g.CountPlayers() > 6 {
		return ErrTooManyPlayers
	}
	for i := 0; i < numberOfPlayers; i++ {
		g.Players[i] = newPlayer()
	}
	return nil
}

func (g *Game) Deal() {
	//Burn 4
	g.Deck.Draw(4)
	//Deal Regular Players
	for _, player := range g.Players {
		cards := g.Deck.Draw(7)
		if player != nil {
			player.ReceiveCards(cards)
		}
	}
	//Deal Dealers Cards
	cards := g.Deck.Draw(7)
	g.Dealer.ReceiveCards(cards)

}

func (g *Game) PlayDealerHand() {
	handslice := g.Dealer.Hand[:]
	hands := FindBestHands(&handslice)
	g.Dealer.HighHand = hands.HighHand
	g.Dealer.LowHand = hands.LowHand
}

type Hands struct {
	LowHand  [2]*Card
	HighHand [5]*Card
}

func FindBestHands(hand *[]*Card) Hands {
	bestHand := Hands{}
	bestHighHand := extractBestHand(hand)
	bestLowHand := extractBestHand(hand)
	// Fill remaining cards if bestLowHand is less than 2 cards
	if len(bestLowHand) < 2 {
		remainingCards := append([]*Card{}, *hand...)
		for _, card := range bestHighHand {
			remainingCards = removeCard(remainingCards, card)
		}
		for _, card := range bestLowHand {
			remainingCards = removeCard(remainingCards, card)
		}
		for len(bestLowHand) < 2 && len(remainingCards) > 0 {
			bestLowHand = append(bestLowHand, remainingCards[0])
			remainingCards = remainingCards[1:]
		}
	}

	// Fill remaining cards if bestHighHand is less than 5 cards
	if len(bestHighHand) < 5 {
		remainingCards := append([]*Card{}, *hand...)
		for _, card := range bestHighHand {
			remainingCards = removeCard(remainingCards, card)
		}
		for _, card := range bestLowHand {
			remainingCards = removeCard(remainingCards, card)
		}
		for len(bestHighHand) < 5 && len(remainingCards) > 0 {
			bestHighHand = append(bestHighHand, remainingCards[0])
			remainingCards = remainingCards[1:]
		}
	}

	bestHand.HighHand = convertToFiveCardArray(bestHighHand)
	bestHand.LowHand = convertToTwoCardArray(bestLowHand)
	return bestHand
}
