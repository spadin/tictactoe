package game

import (
	"fmt"
	"github.com/remogatto/prettytest"
	"testing"
)

type testTictactoeBoardSuite struct {
	prettytest.Suite
}

func TestTictactoeBoardRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testTictactoeBoardSuite),
	)
}

func tictactoeBoardSuiteSetup() *TictactoeBoard {
	return NewTictactoeBoard()
}

func (t *testTictactoeBoardSuite) TestEmptyIndexForNewTictactoeBoard() {
	board := tictactoeBoardSuiteSetup()
	t.True(board.IsEmpty(0), "new tictactoeBoard should have empty position")
}

func (t *testTictactoeBoardSuite) TestNonEmptyTictactoeBoardIndex() {
	board := tictactoeBoardSuiteSetup()
	board.state[0] = X
	t.False(board.IsEmpty(0), "marked position not empty")
}

func (t *testTictactoeBoardSuite) TestMarkingTheTictactoeBoard() {
	board := tictactoeBoardSuiteSetup()
	board.SetMark(X, 0)
	t.False(board.IsEmpty(0), "marks tictactoeBoard a index 0")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardIsFull() {
	board := tictactoeBoardSuiteSetup()
	PlayTiedGame(board)
	t.True(board.IsFull(), "determines the tictactoeBoard is full")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardIsNotFull() {
	board := tictactoeBoardSuiteSetup()
	t.False(board.IsFull(), "determines tictactoeBoard is not full")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardOpenPositions() {
	board := tictactoeBoardSuiteSetup()
	MakeMoves(board, 0, 1, 2)

	expected, actual := []int{3, 4, 5, 6, 7, 8}, board.OpenPositions()

	strExpected := fmt.Sprintf("%v", expected)
	strActual := fmt.Sprintf("%v", actual)

	t.Equal(strExpected, strActual, "returns open positions")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardStringRepresentation() {
	board := tictactoeBoardSuiteSetup()

	vertBars := "     |     |     \n"
	horzBars := "-----------------\n"
	expected := "\n" + vertBars +
		"  0  |  1  |  2  \n" +
		vertBars + horzBars + vertBars +
		"  3  |  4  |  5  \n" +
		vertBars + horzBars + vertBars +
		"  6  |  7  |  8  \n" +
		vertBars

	t.Equal(expected, board.String(), "returns an accurate representation of the board")
}

func (t *testTictactoeBoardSuite) TestSetMarkAddsToHistory() {
	board := tictactoeBoardSuiteSetup()
	board.SetMark(X, 0)
	board.SetMark(O, 5)
	t.Equal(5, board.history[1])
}

func (t *testTictactoeBoardSuite) TestUndo() {
	board := tictactoeBoardSuiteSetup()
	board.SetMark(X, 0)
	board.SetMark(O, 5)
	t.Equal([9]Mark{X, E, E, E, E, O, E, E, E}, board.state)

	board.Undo()
	t.Equal([9]Mark{X, E, E, E, E, E, E, E, E}, board.state)
}
