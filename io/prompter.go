package io

type Prompter interface {
	PromptInt(message string) (input int)
	PromptChoiceList(message string, choices ...string) (selection string)
	PromptIntChoice(message string, choices ...int) (selection int)
}
