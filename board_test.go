package tictactoe

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestEmptyBoardIndexForNewBoard(t *testing.T) {
	board := new(Board)
	assert.True(t, board.IsEmpty(0), "board should have empty index 0")
}

func TestNonEmptyBoardIndex(t *testing.T) {
	board := new(Board)
	board[0] = X
	assert.False(t, board.IsEmpty(0), "marked position not empty")
}

func TestMarkingTheBoard(t *testing.T) {
	board := new(Board)
	board.SetMark(X, 0)
	assert.False(t, board.IsEmpty(0), "marks board a index 0")
}
