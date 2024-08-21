package main

import (
	"fmt"
	"os"
	//"github.com/fatih/color"
	"paigow/pkg/game"

	"github.com/nsf/termbox-go"
)

const (
	HeartsIcon   = "♥"
	DiamondsIcon = "♦"
	ClubsIcon    = "♣"
	SpadesIcon   = "♠"
)

type GameState int

const (
	StateBetting GameState = iota
	StateCardSelection
	StateReveal
	StatePayout
	StateGameOver
)

func printCard(card *game.Card) {
	valueStr := game.ValueToString(card.Value)
	suitStr := game.SuitToString(card.Suit)

	if card.Value == 10 {
		valueStr = "10"
	}

	fmt.Printf(" _____ \n")
	fmt.Printf("|%s%4s|\n", valueStr, " ")
	fmt.Printf("| %s   |\n", suitStr)
	fmt.Printf("|     |\n")
	fmt.Printf("| %s   |\n", suitStr)
	fmt.Printf("|%4s%s|\n", " ", valueStr)
	fmt.Printf("|_____| \n\n")
}

func getCardAscii(card *game.Card) []string {
	var topRow, middleRow, middleRow2, middleRow3, bottomRow, bottomRow2 string
	if card != nil {
		valueStr := game.ValueToString(card.Value)
		suitStr := game.SuitToString(card.Suit)

		if card.Value == 10 {
			valueStr = "10"
		}

		topRow = " _____ "
		middleRow = fmt.Sprintf("|%s%4s|", valueStr, " ")
		middleRow2 = fmt.Sprintf("| %s |", suitStr)
		middleRow3 = fmt.Sprintf("| %s |", suitStr)
		bottomRow = fmt.Sprintf("|%4s%s|", " ", valueStr)
		bottomRow2 = "|_____|"

	}

	handBuffer := []string{
		topRow,
		middleRow,
		middleRow2,
		middleRow3,
		bottomRow,
		bottomRow2,
	}
	return handBuffer
}

func renderHand(hand []*game.Card, hover int) {
	for index, card := range hand {
		var handBuffer []string
		if card != nil {
			handBuffer = getCardAscii(card)
		} else {
			handBuffer = getEmptyCardAscii()
		}
		for y, line := range handBuffer {
			color := termbox.ColorWhite
			if index == hover {
				color = termbox.ColorGreen
			}
			printText(index*len(line), y, color, line)
		}
	}
	termbox.Flush()
}

func getEmptyCardAscii() []string {
	var topRow, middleRow, middleRow2, middleRow3, bottomRow, bottomRow2 string
	topRow = " _____ "
	middleRow = fmt.Sprintf("|%s%4s|", " ", " ")
	middleRow2 = fmt.Sprintf("| %s |", " ")
	middleRow3 = fmt.Sprintf("| %s |", " ")
	bottomRow = fmt.Sprintf("|%4s%s|", " ", " ")
	bottomRow2 = "|_____|"

	handBuffer := []string{
		topRow,
		middleRow,
		middleRow2,
		middleRow3,
		bottomRow,
		bottomRow2,
	}
	return handBuffer

}

func renderEmptyCard() {
}

func printText(x, y int, color termbox.Attribute, text string) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, color, termbox.ColorDefault)
	}
}

type GameView struct {
	Hover    int
	Selected []int
	State    GameState
	Game     *game.Game
}

func handleCardSelection(event termbox.Event, view *GameView, hand []*game.Card) {
	switch event.Key {
	case termbox.KeyArrowLeft:
		if view.Hover > 0 {
			view.Hover--
			renderHand(hand, view.Hover)
		}
	case termbox.KeyArrowRight:
		if view.Hover < len(hand)-1 {
			view.Hover++
			renderHand(hand, view.Hover)
		}
	case termbox.KeyEnter:
		view.Selected = append(view.Selected, view.Hover)
		src := (view.Game.Players[0].Hand)[:]
		dest := view.Game.Players[0].HighHand[:]
		view.Game.Players[0].MoveCards(&src,
			&dest, []int{view.Hover})
		// Handle selection logic here
		renderHand(hand, view.Hover)
	case termbox.KeyEsc:
		view.State = StateGameOver
		os.Exit(1)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termbox.Close()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	newGame := game.NewGame()
	newGame.AddPlayers(1)
	newGame.NewRound()
	view := GameView{}
	view.Game = newGame
	renderHand(newGame.Players[0].Hand[:], 0)
	view.State = StateCardSelection
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			switch view.State {
			case StateCardSelection:
				handleCardSelection(event, &view, newGame.Players[0].Hand[:])
			}

			if event.Key == termbox.KeyEsc {
				break
			}
		}
	}
}
