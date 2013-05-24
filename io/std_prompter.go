package io

import (
	"strconv"
	"strings"
)

type StdPrompter struct {
	Getter
	Printer
}

func NewStdPrompter() (prompter *StdPrompter) {
	prompter = &StdPrompter{Getter: NewStdinGetter(), Printer: NewStdoutPrinter()}
	return
}

func (p *StdPrompter) PromptInt(message string) (input int) {
	promptMessage := message + ": "
	p.Print(promptMessage)
	input = p.GetIntWithCallback(func() {
		p.Println("Invalid input, please enter a number")
		p.Print(promptMessage)
	})
	return
}

func (p *StdPrompter) PromptChoiceList(message string, choices ...string) (selection string) {
	p.Println(message)
	for index, choice := range choices {
		choiceStr := "" + strconv.Itoa(index+1) + ". " + choice
		p.Println(choiceStr)
	}

	selectionInt := p.PromptInt("Enter your choice")
	selectionIdx := selectionInt - 1
	for selectionIdx > len(choices) || selectionIdx < 0 {
		p.Println("Invalid input, please enter a number from the choice list")
		selectionInt = p.PromptInt("Enter your choice")
		selectionIdx = selectionInt - 1

	}
	selection = choices[selectionIdx]
	return
}

func validSelection(selection int, choices []int) (validSelection bool) {
	validSelection = false
	for _, choice := range choices {
		if selection == choice {
			validSelection = true
		}
	}
	return
}

func intArrayToStringArray(intArr []int) (strArr []string) {
	for _, item := range intArr {
		strArr = append(strArr, strconv.Itoa(item))
	}
	return
}

func (p *StdPrompter) PromptIntChoice(message string, choices ...int) (selection int) {
	stringChoices := intArrayToStringArray(choices)

	promptMessage := message + " (" + strings.Join(stringChoices, ", ") + ")"
	selection = p.PromptInt(promptMessage)

	for !validSelection(selection, choices) {
		selection = p.PromptInt("Invalid input, please choose one of the numbers on the list\n" +
			promptMessage)
	}

	return
}
