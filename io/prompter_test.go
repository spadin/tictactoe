package io

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/io/helpers"
	"testing"
)

type testPrompterSuite struct {
	prettytest.Suite
}

func prompterSuiteSetup() (prompter *Prompter) {
	prompter = &Prompter{in: NewStdinGetter(), out: NewStdoutPrinter()}
	return
}

func TestPrompterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testPrompterSuite),
	)
}

func (t *testPrompterSuite) TestPromptInt() {
	var input int
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateInput("1", func() {
			prompter := prompterSuiteSetup()
			input = prompter.PromptInt("Enter a number:")
		})
	})

	t.Equal("Enter a number:", output)
	t.Equal(1, input)
}

func (t *testPrompterSuite) TestPromptIntRepeatedly() {
	var input int
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateMultipleInput([]string{"A", "1"}, func() {
			prompter := prompterSuiteSetup()
			input = prompter.PromptInt("Enter a number:")
		})
	})

	t.Equal("Enter a number:\nInvalid input, please enter a number\nEnter a number:", output)
	t.Equal(1, input)
}
