package main

import (
	"fmt"
	"os"
	//"github.com/fatih/color"
	"paigow/pkg/game"

	"github.com/nsf/termbox-go"
)

const (
	HeartsIcon   = "H"
	DiamondsIcon = "D"
	ClubsIcon    = "C"
	SpadesIcon   = "S"
)

type GameState int

const (
	StateBetting GameState = iota
	StateCardSelection
	StateReveal
	StatePayout
	StateGameOver
)

const (
	mainHandY = 0
	highHandY = 10
	lowHandY  = 20
)

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
		middleRow2 = fmt.Sprintf("|  %s  |", suitStr)
		middleRow3 = fmt.Sprintf("|  %s  |", suitStr)
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

func renderHand(hand []*game.Card, hover, startY int) {
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
			printText(index*len(line), y+startY, color, line)
		}
	}
	termbox.Flush()
}

func renderAllHands(player *game.Player, hover int) {
	//Render Player Hand
	renderHand(player.Hand[:], hover, mainHandY)
	//Render High Hand
	renderHand(player.HighHand[:], hover-7, highHandY)
	//Render Low Hand
	renderHand(player.LowHand[:], hover-12, lowHandY)
}

func getEmptyCardAscii() []string {
	var topRow, middleRow, middleRow2, middleRow3, bottomRow, bottomRow2 string
	topRow = " _____ "
	middleRow = fmt.Sprintf("|%5s|", " ")
	middleRow2 = fmt.Sprintf("| %3s |", " ")
	middleRow3 = fmt.Sprintf("| %3s |", " ")
	bottomRow = fmt.Sprintf("|%5s|", " ")
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
			renderAllHands(view.Game.Players[0], view.Hover)
		}
	case termbox.KeyArrowRight:
		if view.Hover < 14-1 {
			view.Hover++
			renderAllHands(view.Game.Players[0], view.Hover)
		}
	case termbox.KeyEnter:
		view.Selected = append(view.Selected, view.Hover)
		src := (view.Game.Players[0].Hand)[:]
		dest := view.Game.Players[0].HighHand[:]
		view.Game.Players[0].MoveCards(&src,
			&dest, []int{view.Hover})
		// Handle selection logic here
		renderAllHands(view.Game.Players[0], view.Hover)
	case termbox.KeyArrowDown:
		if view.Hover <= 6 {
			view.Hover += 7
		} else if view.Hover <= 11 {
			view.Hover += 5
		} else {
			//Break Since no re-render required
			break
		}
		renderAllHands(view.Game.Players[0], view.Hover)

	case termbox.KeyArrowUp:
		if view.Hover >= 12 {
			view.Hover -= 5
		} else if view.Hover >= 7 {
			view.Hover -= 7
		} else {
			//Break Since no re-render required
			break
		}
		renderAllHands(view.Game.Players[0], view.Hover)

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
	renderHand(newGame.Players[0].Hand[:], 0, mainHandY)
	renderHand(newGame.Players[0].HighHand[:], -1, highHandY)
	renderHand(newGame.Players[0].LowHand[:], -1, lowHandY)
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
