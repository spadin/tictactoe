package io

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/io/helpers"
	"os"
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

func (t *testPrompterSuite) TestNewPrompter() {
	prompter := NewPrompter()
	t.Equal(os.Stdin, prompter.in.in, "sets Stdin to getter")
	t.Equal(os.Stdout, prompter.out.out, "sets Stdout to printer")

}

func (t *testPrompterSuite) TestPromptChoice() {
	var choice string
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateInput("1", func() {
			prompter := prompterSuiteSetup()
			choice = prompter.PromptChoice("Choose from the following", "banana", "apple")
		})
	})

	t.Equal("Choose from the following\n1. banana\n2. apple\nEnter your choice:", output)
	t.Equal("banana", choice, "should choose choice 1 and return choice string")
}

func (t *testPrompterSuite) TestPromptChoiceWithInvalidInputs() {
	var choice string
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateMultipleInput([]string{"5", "1"}, func() {
			prompter := prompterSuiteSetup()
			choice = prompter.PromptChoice("Choose from the following", "banana", "apple")
		})
	})

	t.Equal("Choose from the following\n1. banana\n2. apple\nEnter your choice:\nInvalid input, please enter a number from the choice list\nEnter your choice:", output)
	t.Equal("banana", choice, "should choose choice 1 and return choice string")
}
