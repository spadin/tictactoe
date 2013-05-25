package player

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
	"testing"
)

type testAiPlayerSuite struct {
	prettytest.Suite
}

func aiPlayerSuiteSetup() (board game.Board, player *Ai) {
	prompter := io.NewStdPrompter()
	board = game.NewTictactoeBoard()
	p := NewAiPlayer(game.X, board, prompter)
	player = p.(*Ai)
	return
}

func TestAiPlayerRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testAiPlayerSuite),
	)
}

func (t *testAiPlayerSuite) TestCalculateScoreForTiedGame() {
	board, player := aiPlayerSuiteSetup()
	game.PlayTiedGame(board)
	t.Equal(0, player.calculateScore())
}

func (t *testAiPlayerSuite) TestCalculateScoreForWinningGame() {
	board, player := aiPlayerSuiteSetup()
	game.PlayXWinningGame(board)
	t.Equal(1, player.calculateScore())
}

func (t *testAiPlayerSuite) TestCalculateScoreForLosingGame() {
	board, player := aiPlayerSuiteSetup()
	game.PlayOWinningGame(board)
	t.Equal(-1, player.calculateScore())
}

func (t *testAiPlayerSuite) TestPlayGameTyingMove() {
	board, player := aiPlayerSuiteSetup()
	game.MakeMoves(board, 0, 1, 2, 4, 3, 5, 7, 6)
	move := player.Play()
	t.Equal(8, move)
}

func (t *testAiPlayerSuite) TestPlayWinningMove() {
	board, player := aiPlayerSuiteSetup()
	game.MakeMoves(board, 0, 4, 1, 5)
	move := player.Play()
	t.Equal(2, move)
}

func (t *testAiPlayerSuite) TestPlayBestInitialMove() {
	_, player := aiPlayerSuiteSetup()
	move := player.Play()
	t.Equal(0, move)
}
