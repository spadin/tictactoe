package tictactoe

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testBoardSuite struct {
	prettytest.Suite
}

func TestBoardRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testBoardSuite),
	)
}

func setup() *Board {
	return new(Board)
}

func (t *testBoardSuite) TestEmptyBoardIndexForNewBoard() {
	board := setup()
	t.True(board.IsEmpty(0), "new board should have empty position")
}

func (t *testBoardSuite) TestNonEmptyBoardIndex() {
	board := setup()
	board[0] = X
	t.False(board.IsEmpty(0), "marked position not empty")
}

func (t *testBoardSuite) TestMarkingTheBoard() {
	board := setup()
	board.SetMark(X, 0)
	t.False(board.IsEmpty(0), "marks board a index 0")
}
