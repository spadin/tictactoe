package tictactoe

import (
	"github.com/remogatto/prettytest"
	"os"
	"testing"
)

type testGetterSuite struct {
	prettytest.Suite
}

func getterSuiteSetup() (getter *Getter) {
	getter = &Getter{in: os.Stdin}
	return
}

func TestGetterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testGetterSuite),
	)
}

func (t *testGetterSuite) TestFscan() {
	var input string
	simulateInput("sample-input", func() {
		getter := getterSuiteSetup()
		input = getter.Fscan()
	})

	t.Equal("sample-input", input)
}

func (t *testGetterSuite) TestNewGetter() {
	_, reader, _ := os.Pipe()
	getter := NewGetter(reader)

	t.Equal(reader, getter.in)
}

func (t *testGetterSuite) TestNewStdinGetter() {
	getter := NewStdinGetter()
	t.Equal(os.Stdin, getter.in)
}
