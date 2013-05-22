package game

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testTictactoeGameSuite struct {
	prettytest.Suite
}

func gameSuiteSetup() (board Board, game *TictactoeGame) {
	board = NewTictactoeBoard()
	game = NewTictactoeGame(board)
	return
}

func TestTictactoeGameRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testTictactoeGameSuite),
	)
}

func (t *testTictactoeGameSuite) TestNewTictactoeGameHasNoWinner() {
	_, game := gameSuiteSetup()
	t.False(game.HasWinner(), "New game has no winner")
}

func (t *testTictactoeGameSuite) TestXWinningTictactoeGameHasWinner() {
	board, game := gameSuiteSetup()
	playXWinningGame(board.(*TictactoeBoard))
	t.True(game.HasWinner(), "X winning game should have winner")
}

func (t *testTictactoeGameSuite) TestTictactoeGameoverFalseWithNewGame() {
	_, game := gameSuiteSetup()
	t.False(game.IsGameover(), "New game is not gameover")
}

func (t *testTictactoeGameSuite) TestTictactoeGameoverTrueWithWinningGame() {
	board, game := gameSuiteSetup()
	playXWinningGame(board.(*TictactoeBoard))
	t.True(game.IsGameover(), "TictactoeGame is over when there is a winner")
}

func (t *testTictactoeGameSuite) TestTictactoeGameoverTrueWithFullBoard() {
	board, game := gameSuiteSetup()
	playTiedGame(board.(*TictactoeBoard))
	t.True(game.IsGameover(), "TictactoeGame is over when board full")
}
