package game

type Player struct {
	Hand         [7]*Card
	LowHand      [2]*Card
	HighHand     [5]*Card
	HighHandRank *HandRank
	LowHandRank  *HandRank
	Balance      int
	MainBet      int
}

type PlayerErr string

func (e PlayerErr) Error() string {
	return string(e)
}

const (
	ErrNotEnoughMoney = PlayerErr("Could not make bet, Not enough money")
	ErrNoSpaceInHand  = PlayerErr("Could not move Card, no Space in destination")
	ErrNoCardsInHand  = PlayerErr("Cold not move Card, no cards in source hand")
)

func newPlayer() *Player {
	newPlayer := Player{Balance: 0, MainBet: 0}
	return &newPlayer
}

func (p *Player) Bet(bet int) error {
	if bet > p.Balance {
		return ErrNotEnoughMoney
	}
	p.MainBet = bet
	p.Balance -= bet
	return nil
}

func (p *Player) ReceiveCards(cards []*Card) {
	for i := 0; i < len(p.Hand); i++ {
		p.Hand[i] = cards[i]
	}
}

func (p *Player) MoveCards(src *[]*Card, dest *[]*Card, indexes []int) error {
	numberOfCardsToMove := len(indexes)
	emptySpaceCount := 0
	for _, card := range *dest {
		if card == nil {
			emptySpaceCount = emptySpaceCount + 1
		}

	}
	if emptySpaceCount < numberOfCardsToMove {
		return ErrNoSpaceInHand
	}
	for _, potentialIndex := range indexes {
		if (*src)[potentialIndex] == nil {
			return ErrNoCardsInHand
		} else {
			for _, card := range *dest {
				if card == nil {
					card = (*src)[potentialIndex]
					(*src)[potentialIndex] = nil
				}
			}
		}
	}
	return nil
}
