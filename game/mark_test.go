package game

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testMarkSuite struct {
	prettytest.Suite
}

func TestMarkRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testMarkSuite),
	)
}

func (t *testMarkSuite) TestXMarkIsOne() {
	t.Equal(X, Mark(1), "X should equal to 1")
}

func (t *testMarkSuite) TestEMarkIsZero() {
	t.Equal(E, Mark(0), "E should equal 0")
}

func (t *testMarkSuite) TestOMarkIsTwo() {
	t.Equal(O, Mark(2), "O should equal 2")
}

func (t *testMarkSuite) TestStringifyX() {
	t.Equal("X", stringify(X), "X mark should be X string")
}

func (t *testMarkSuite) TestStringifyO() {
	t.Equal("O", stringify(O), "O mark should be O string")
}

func (t *testMarkSuite) TestStringifyEmptyMark() {
	t.Equal(" ", stringify(E), "empty mark should be empty space")
}
