package player

import (
	"fmt"
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
	"testing"
)

type testPlayersSuite struct {
	prettytest.Suite
}

func playersSuiteSetup() (players *TictactoePlayers) {
	players = &TictactoePlayers{}
	return
}

func TestPlayersRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testPlayersSuite),
	)
}

func (t *testPlayersSuite) TestAddPlayer() {
	players := playersSuiteSetup()

	human := NewHumanPlayer(game.X, game.NewTictactoeBoard(), io.NewStdPrompter())
	players.Add(human)
	t.Equal(1, len(players.list))

	human2 := NewHumanPlayer(game.X, game.NewTictactoeBoard(), io.NewStdPrompter())
	players.Add(human2)
	t.Equal(2, len(players.list))

	human3 := NewHumanPlayer(game.X, game.NewTictactoeBoard(), io.NewStdPrompter())
	err := players.Add(human3)
	t.Equal(fmt.Sprint(err), "Can only have two players")
	t.Equal(2, len(players.list))
}

func (t *testPlayersSuite) TestNextPlayer() {
	players := playersSuiteSetup()

	player1 := NewHumanPlayer(game.X, game.NewTictactoeBoard(), io.NewStdPrompter())
	player2 := NewHumanPlayer(game.O, game.NewTictactoeBoard(), io.NewStdPrompter())
	players.Add(player1)
	players.Add(player2)

	t.Equal(player2, players.Next(player1))
	t.Equal(player1, players.Next(player2))
	t.Equal(player1, players.Next(nil))
}

func (t *testPlayersSuite) TestAvailablePlayerTypeChoices() {
	players := playersSuiteSetup()
	actual := players.AvailablePlayerTypeChoices()
	t.Equal("human", actual[0])
}

func (t *testPlayersSuite) TestMake() {
	players := playersSuiteSetup()
	mark := game.X
	board := game.NewTictactoeBoard()
	prompter := io.NewStdPrompter()

	player := players.Make("human", mark, board, prompter)
	_, err := player.(*Human)
	t.Nil(err, "should creating human player")
}

func (t *testPlayersSuite) TestMakeAndAdd() {
	players := playersSuiteSetup()
	mark := game.X
	board := game.NewTictactoeBoard()
	prompter := io.NewStdPrompter()

	player := players.MakeAndAdd("human", mark, board, prompter)
	t.Equal(1, len(players.list))
	_, err := player.(*Human)
	t.Nil(err, "should creating human player")
}
