package player

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
	"testing"
)

const X, O, E = game.X, game.O, game.E

type testPlayerSuite struct {
	prettytest.Suite
}

func playerSuiteSetup() (board game.Board, player Player) {
	board = game.NewTictactoeBoard()
	prompter := io.NewStdPrompter()
	player = &Human{board: board, prompter: prompter, Mark: game.X}
	return
}

func TestPlayerRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testPlayerSuite),
	)
}

func (t *testPlayerSuite) TestPlay() {
	var choice int
	io.CaptureOutput(func() {
		io.SimulateInput("4", func() {
			_, player := playerSuiteSetup()
			choice = player.Play()
		})
	})

	t.Equal(4, choice, "player should have chosen 4")
}
