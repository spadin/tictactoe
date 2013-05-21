package io

import (
	"strconv"
	"strings"
)

type Prompter struct {
	in  *Getter
	out *Printer
}

func (p *Prompter) PromptInt(message string) (input int) {
	p.out.Print(message)
	input = p.in.GetIntWithCallback(func() {
		p.out.Println("\nInvalid input, please enter a number")
		p.out.Print(message)
	})
	return
}

func (p *Prompter) PromptChoice(message string, choices ...string) (selection string) {
	p.out.Println(message)
	for index, choice := range choices {
		choiceStr := strings.Join([]string{strconv.Itoa(index + 1), ". ", choice}, "")
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

func NewPrompter() (prompter *Prompter) {
	prompter = &Prompter{in: NewStdinGetter(), out: NewStdoutPrinter()}
	return
}
