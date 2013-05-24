package game

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testGameHelpersSuite struct {
	prettytest.Suite
}

func TestGameHelpersRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testGameHelpersSuite),
	)
}

func (t *testGameHelpersSuite) TestIsEven() {
	t.True(isEven(0))
	t.False(isEven(1))
	t.True(isEven(2))
}

func (t *testGameHelpersSuite) TestMakeMoves() {
	board := NewTictactoeBoard()
	makeMoves(board, 0, 1)
	t.Equal(X, board[0])
	t.Equal(O, board[1])
}

func (t *testGameHelpersSuite) TestPlayXWinningGame() {
	board := NewTictactoeBoard()
	game := NewTictactoeRules(board)
	playXWinningGame(board)
	t.True(game.IsGameover())
	t.Equal(X, game.Winner())
}

func (t *testGameHelpersSuite) TestPlayTiedGame() {
	board := NewTictactoeBoard()
	game := NewTictactoeRules(board)
	playTiedGame(board)
	t.True(game.IsGameover())
	t.Nil(game.Winner())
}
