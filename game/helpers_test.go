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
	t.True(IsEven(0))
	t.False(IsEven(1))
	t.True(IsEven(2))
}

func (t *testGameHelpersSuite) TestMakeMoves() {
	board := NewTictactoeBoard()
	MakeMoves(board, 0, 1)
	t.Equal(X, board.state[0])
	t.Equal(O, board.state[1])
}

func (t *testGameHelpersSuite) TestPlayXWinningGame() {
	board := NewTictactoeBoard()
	game := NewTictactoeRules(board)
	PlayXWinningGame(board)
	t.True(game.IsGameover())
	t.Equal(X, game.Winner())
}

func (t *testGameHelpersSuite) TestPlayOWinningGame() {
	board := NewTictactoeBoard()
	game := NewTictactoeRules(board)
	PlayOWinningGame(board)
	t.True(game.IsGameover())
	t.Equal(O, game.Winner())
}

func (t *testGameHelpersSuite) TestPlayTiedGame() {
	board := NewTictactoeBoard()
	game := NewTictactoeRules(board)
	PlayTiedGame(board)
	t.True(game.IsGameover())
	t.Nil(game.Winner())
}
