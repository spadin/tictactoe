package io

import (
	"strconv"
	"strings"
)

type StdPrompter struct {
	in  Getter
	out Printer
}

func NewStdPrompter() (prompter *StdPrompter) {
	prompter = &StdPrompter{in: NewStdinGetter(), out: NewStdoutPrinter()}
	return
}

func (p *StdPrompter) PromptInt(message string) (input int) {
	p.out.Print(message)
	input = p.in.GetIntWithCallback(func() {
		p.out.Println("\nInvalid input, please enter a number")
		p.out.Print(message)
	})
	return
}

func (p *StdPrompter) PromptChoiceList(message string, choices ...string) (selection string) {
	p.out.Println(message)
	for index, choice := range choices {
		choiceStr := "" + strconv.Itoa(index+1) + ". " + choice
		p.out.Println(choiceStr)
	}

	selectionInt := p.PromptInt("Enter your choice:")
	selectionIdx := selectionInt - 1
	for selectionIdx > len(choices) || selectionIdx < 0 {
		p.out.Println("\nInvalid input, please enter a number from the choice list")
		selectionInt = p.PromptInt("Enter your choice:")
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

	promptMessage := message + ": (" + strings.Join(stringChoices, ", ") + ")"
	selection = p.PromptInt(promptMessage)

	for !validSelection(selection, choices) {
		selection = p.PromptInt("\nInvalid input, please choose one of the numbers listed above")
	}

	return
}
