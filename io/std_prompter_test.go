package io

import (
	"github.com/remogatto/prettytest"
	"testing"
)

type testStdPrompterSuite struct {
	prettytest.Suite
}

func stdPrompterSuiteSetup() (stdPrompter *StdPrompter) {
	stdPrompter = &StdPrompter{in: NewStdinGetter(), out: NewStdoutPrinter()}
	return
}

func TestStdPrompterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testStdPrompterSuite),
	)
}

func (t *testStdPrompterSuite) TestPromptInt() {
	var input int
	output := CaptureOutput(func() {
		SimulateInput("1", func() {
			stdPrompter := stdPrompterSuiteSetup()
			input = stdPrompter.PromptInt("Enter a number:")
		})
	})

	t.Equal("Enter a number:", output)
	t.Equal(1, input)
}

func (t *testStdPrompterSuite) TestPromptIntRepeatedly() {
	var input int
	output := CaptureOutput(func() {
		SimulateMultipleInput([]string{"A", "1"}, func() {
			stdPrompter := stdPrompterSuiteSetup()
			input = stdPrompter.PromptInt("Enter a number:")
		})
	})

	expectedOutput := "Enter a number:\n" +
		"Invalid input, please enter a number\n" +
		"Enter a number:"

	t.Equal(expectedOutput, output)
	t.Equal(1, input)
}

func (t *testStdPrompterSuite) TestNewStdPrompter() {
	stdPrompter := NewStdPrompter()
	_, getterErr := stdPrompter.in.(*StdinGetter)
	t.Nil(getterErr, "sets in to StdinGetter")

	_, printerErr := stdPrompter.out.(*StdoutPrinter)
	t.Nil(printerErr, "sets out to StdoutPrinter")
}

func (t *testStdPrompterSuite) TestPromptChoiceList() {
	var choice string
	output := CaptureOutput(func() {
		SimulateInput("1", func() {
			stdPrompter := stdPrompterSuiteSetup()
			choice = stdPrompter.PromptChoiceList("Choose from the following", "banana", "apple")
		})
	})

	expectedOutput := "Choose from the following\n" +
		"1. banana\n" +
		"2. apple\n" +
		"Enter your choice:"

	t.Equal(expectedOutput, output)
	t.Equal("banana", choice, "should choose choice 1 and return choice string")
}

func (t *testStdPrompterSuite) TestPromptChoiceListWithInvalidInputs() {
	var choice string
	output := CaptureOutput(func() {
		SimulateMultipleInput([]string{"5", "1"}, func() {
			stdPrompter := stdPrompterSuiteSetup()
			choice = stdPrompter.PromptChoiceList("Choose from the following", "banana", "apple")
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

func (t *testStdPrompterSuite) TestPromptIntChoice() {
	var choice int
	output := CaptureOutput(func() {
		SimulateInput("6", func() {
			stdPrompter := stdPrompterSuiteSetup()
			choice = stdPrompter.PromptIntChoice("Choose from the following", 1, 2, 5, 6)
		})
	})

	expectedOutput := "Choose from the following: (1, 2, 5, 6)"

	t.Equal(expectedOutput, output)
	t.Equal(6, choice, "should choose 6")
}

func (t *testStdPrompterSuite) TestPromptIntChoiceWithInvalidInput() {
	var choice int
	output := CaptureOutput(func() {
		SimulateMultipleInput([]string{"8", "3", "5"}, func() {
			stdPrompter := stdPrompterSuiteSetup()
			choice = stdPrompter.PromptIntChoice("Choose from the following", 1, 2, 5, 6)
		})
	})

	expectedOutput := "Choose from the following: (1, 2, 5, 6)\n" +
		"Invalid input, please choose one of the numbers listed above\n" +
		"Invalid input, please choose one of the numbers listed above"

	t.Equal(expectedOutput, output)
	t.Equal(5, choice, "should choose 5")
}
