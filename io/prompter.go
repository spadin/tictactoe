package io

type Prompter interface {
	PromptInt(string) int
	PromptChoiceList(string, ...string) string
	PromptIntChoice(string, ...int) int
}
