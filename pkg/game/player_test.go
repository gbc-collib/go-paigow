package game

import "testing"

func TestBet(t *testing.T) {
	player := Player{Balance: 0, MainBet: 0}

	t.Run("Bet without Enough Money", func(t *testing.T) {
		err := player.Bet(10)
		assertError(t, err, ErrNotEnoughMoney)

	})

	t.Run("Bet With enough Money", func(t *testing.T) {
		player := Player{Balance: 10, MainBet: 0}
		err := player.Bet(10)
		if err != nil {
			t.Errorf("Got error, when no error expected")
		}
		wantedBalance := 0
		if player.Balance != 0 {
			t.Errorf("Got %d, wanted %d", player.Balance, wantedBalance)
		}
		wantedBet := 10
		if player.MainBet != wantedBet {
			t.Errorf("Got %d, wanted %d", player.MainBet, wantedBet)
		}
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func testMoveCards(t *testing.T) {
	src := []*Card{}
	dest := []*Card{}
	index := 0
    //Should Move Cards

    //Should Return Error fr nil src


    //Should return err for no dest space

}
