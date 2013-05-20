package tictactoe

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestXMarkIsOne(t *testing.T) {
	assert.Equal(t, X, 1, "X should equal to 1")
}

func TestOMarkIsTwo(t *testing.T) {
	assert.Equal(t, O, 2, "O should equal 2")
}
