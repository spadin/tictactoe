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
	getter = NewStdinGetter()
	return
}

func TestGetterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testGetterSuite),
	)
}

func (t *testGetterSuite) TestFscanln() {
	var input string
	simulateInput("sample-input", func() {
		getter := getterSuiteSetup()
		input = getter.Fscanln()
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

func (t *testGetterSuite) TestGetString() {
	var input string
	simulateInput("sample-string", func() {
		getter := getterSuiteSetup()
		input = getter.GetString()
	})

	t.Equal("sample-string", input)
}

func (t *testGetterSuite) TestGetInt() {
	var input int
	simulateInput("1", func() {
		getter := getterSuiteSetup()
		input = getter.GetInt()
	})

	t.Equal(1, input)
}

func (t *testGetterSuite) TestRepeatedlyAskForInt() {
	var input int
	simulateMultipleInput([]string{"A", "C", "8"}, func() {
		getter := getterSuiteSetup()
		input = getter.GetInt()
	})
	t.Equal(8, input)
}
