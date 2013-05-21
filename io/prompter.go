package io

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
