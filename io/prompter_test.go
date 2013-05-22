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

	expectedOutput := "Enter a number:\n" +
		"Invalid input, please enter a number\n" +
		"Enter a number:"

	t.Equal(expectedOutput, output)
	t.Equal(1, input)
}

func (t *testPrompterSuite) TestNewPrompter() {
	prompter := NewPrompter()
	t.Equal(os.Stdin, prompter.in.in, "sets Stdin to getter")
	t.Equal(os.Stdout, prompter.out.out, "sets Stdout to printer")

}

func (t *testPrompterSuite) TestPromptChoiceList() {
	var choice string
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateInput("1", func() {
			prompter := prompterSuiteSetup()
			choice = prompter.PromptChoiceList("Choose from the following", "banana", "apple")
		})
	})

	expectedOutput := "Choose from the following\n" +
		"1. banana\n" +
		"2. apple\n" +
		"Enter your choice:"

	t.Equal(expectedOutput, output)
	t.Equal("banana", choice, "should choose choice 1 and return choice string")
}

func (t *testPrompterSuite) TestPromptChoiceListWithInvalidInputs() {
	var choice string
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateMultipleInput([]string{"5", "1"}, func() {
			prompter := prompterSuiteSetup()
			choice = prompter.PromptChoiceList("Choose from the following", "banana", "apple")
		})
	})

	expectedOutput := "Choose from the following\n" +
		"1. banana\n" +
		"2. apple\n" +
		"Enter your choice:\n" +
		"Invalid input, please enter a number from the choice list\n" +
		"Enter your choice:"

	t.Equal(expectedOutput, output)
	t.Equal("banana", choice, "should choose choice 1 and return choice string")
}

func (t *testPrompterSuite) TestPromptIntChoice() {
	var choice int
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateInput("6", func() {
			prompter := prompterSuiteSetup()
			choice = prompter.PromptIntChoice("Choose from the following", 1, 2, 5, 6)
		})
	})

	expectedOutput := "Choose from the following: (1, 2, 5, 6)"

	t.Equal(expectedOutput, output)
	t.Equal(6, choice, "should choose 6")
}

func (t *testPrompterSuite) TestPromptIntChoiceWithInvalidInput() {
	var choice int
	output := iohelpers.CaptureOutput(func() {
		iohelpers.SimulateMultipleInput([]string{"8", "3", "5"}, func() {
			prompter := prompterSuiteSetup()
			choice = prompter.PromptIntChoice("Choose from the following", 1, 2, 5, 6)
		})
	})

	expectedOutput := "Choose from the following: (1, 2, 5, 6)\n" +
		"Invalid input, please choose one of the numbers listed above\n" +
		"Invalid input, please choose one of the numbers listed above"

	t.Equal(expectedOutput, output)
	t.Equal(5, choice, "should choose 5")
}
