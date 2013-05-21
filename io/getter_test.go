package io

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/io/helpers"
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
	iohelpers.SimulateInput("sample-input", func() {
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
	iohelpers.SimulateInput("sample-string", func() {
		getter := getterSuiteSetup()
		input = getter.GetString()
	})

	t.Equal("sample-string", input)
}

func (t *testGetterSuite) TestGetInt() {
	var input int
	iohelpers.SimulateInput("1", func() {
		getter := getterSuiteSetup()
		input = getter.GetInt()
	})

	t.Equal(1, input)
}

func (t *testGetterSuite) TestRepeatedlyAskForInt() {
	var input int
	iohelpers.SimulateMultipleInput([]string{"A", "C", "8"}, func() {
		getter := getterSuiteSetup()
		input = getter.GetInt()
	})
	t.Equal(8, input)
}

func (t *testGetterSuite) TestGetIntWithCallback() {
	var count int
	iohelpers.SimulateMultipleInput([]string{"A", "C", "8"}, func() {
		getter := getterSuiteSetup()
		getter.GetIntWithCallback(func() {
			count += 1
		})
	})

	t.Equal(2, count)
}
