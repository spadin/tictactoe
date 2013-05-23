package game

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testTictactoeMarkSuite struct {
	prettytest.Suite
}

func TestTictactoeMarkRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testTictactoeMarkSuite),
	)
}

func (t *testTictactoeMarkSuite) TestXTictactoeMarkIsOne() {
	t.Equal(X, TictactoeMark(1), "X should equal to 1")
}

func (t *testTictactoeMarkSuite) TestETictactoeMarkIsZero() {
	t.Equal(E, TictactoeMark(0), "E should equal 0")
}

func (t *testTictactoeMarkSuite) TestOTictactoeMarkIsTwo() {
	t.Equal(O, TictactoeMark(2), "O should equal 2")
}

func (t *testTictactoeMarkSuite) TestStringifyX() {
	t.Equal("X", stringify(X), "X mark should be X string")
}

func (t *testTictactoeMarkSuite) TestStringifyO() {
	t.Equal("O", stringify(O), "O mark should be O string")
}

func (t *testTictactoeMarkSuite) TestStringifyEmptyTictactoeMark() {
	t.Equal(" ", stringify(E), "empty mark should be empty space")
}
