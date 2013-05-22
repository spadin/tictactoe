package game

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testGameSuite struct {
	prettytest.Suite
}

func gameSuiteSetup() (board Board, game *Game) {
	board = NewTictactoeBoard()
	game = new(Game)
	game.board = board
	return
}

func TestGameRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testGameSuite),
	)
}

func (t *testGameSuite) TestNewGameHasNoWinner() {
	_, game := gameSuiteSetup()
	t.False(game.HasWinner(), "New game has no winner")
}

func (t *testGameSuite) TestXWinningGameHasWinner() {
	board, game := gameSuiteSetup()
	playXWinningGame(board.(*TictactoeBoard))
	t.True(game.HasWinner(), "X winning game should have winner")
}

func (t *testGameSuite) TestGameoverFalseWithNewGame() {
	_, game := gameSuiteSetup()
	t.False(game.IsGameover(), "New game is not gameover")
}

func (t *testGameSuite) TestGameoverTrueWithWinningGame() {
	board, game := gameSuiteSetup()
	playXWinningGame(board.(*TictactoeBoard))
	t.True(game.IsGameover(), "Game is over when there is a winner")
}

func (t *testGameSuite) TestGameoverTrueWithFullBoard() {
	board, game := gameSuiteSetup()
	playTiedGame(board.(*TictactoeBoard))
	t.True(game.IsGameover(), "Game is over when board full")
}
