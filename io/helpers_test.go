package io

import (
	"fmt"
	"github.com/remogatto/prettytest"
	"testing"
)

type testIoHelpersSuite struct {
	prettytest.Suite
}

func TestIoHelpersRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testIoHelpersSuite),
	)
}

func (t *testIoHelpersSuite) TestSimulateInput() {
	var input string
	SimulateInput("test-input", func() {
		fmt.Scan(&input)
	})
	t.Equal("test-input", input)
}

func (t *testIoHelpersSuite) TestSimulateMultipleInput() {
	var input1 string
	var input2 string
	SimulateMultipleInput([]string{"one", "two"}, func() {
		fmt.Scan(&input1)
		fmt.Scan(&input2)
	})
	t.Equal("one", input1)
	t.Equal("two", input2)
}

func (t *testIoHelpersSuite) TestCaptureOutput() {
	output := CaptureOutput(func() {
		fmt.Print("testing-output")
	})
	t.Equal("testing-output", output)
}
