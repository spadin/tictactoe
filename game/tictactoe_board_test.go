package game

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

func boardSuiteSetup() *Board {
	return NewBoard()
}

func (t *testBoardSuite) TestEmptyBoardIndexForNewBoard() {
	board := boardSuiteSetup()
	t.True(board.IsEmpty(0), "new board should have empty position")
}

func (t *testBoardSuite) TestNonEmptyBoardIndex() {
	board := boardSuiteSetup()
	board[0] = X
	t.False(board.IsEmpty(0), "marked position not empty")
}

func (t *testBoardSuite) TestMarkingTheBoard() {
	board := boardSuiteSetup()
	board.SetMark(X, 0)
	t.False(board.IsEmpty(0), "marks board a index 0")
}

func (t *testBoardSuite) TestBoardIsFull() {
	board := boardSuiteSetup()
	playTiedGame(board)
	t.True(board.IsFull(), "determines the board is full")
}

func (t *testBoardSuite) TestBoardIsNotFull() {
	board := boardSuiteSetup()
	t.False(board.IsFull(), "determines board is not full")
}
