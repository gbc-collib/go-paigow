package game

type Suit int
type Value int

const (
	NoSuit = iota
	Club
	Diamond
	Heart
	Spade
)

const (
	Joker Value = iota
	Two   Value = iota + 1
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Suit  Suit
	Value Value
    Selected bool
}

func ValueToString(value Value) string {
	switch value {
	case Joker:
		return "?"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return "?"
	}
}
func SuitToString(suit Suit) string {
	switch suit {
	case Club:
		return "♣"
	case Diamond:
		return "♦"
	case Heart:
		return "♥"
	case Spade:
		return "♠"
	default:
		return "?"
	}
}

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func (hr HandRank) String() string {
	return [...]string{
		"HighCard",
		"OnePair",
		"TwoPair",
		"ThreeOfAKind",
		"Straight",
		"Flush",
		"FullHouse",
		"FourOfAKind",
		"StraightFlush",
	}[hr]
}
