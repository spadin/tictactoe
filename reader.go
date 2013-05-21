package tictactoe

import (
	"fmt"
	"io"
	"os"
)

type Getter struct {
	in io.Reader
}

func NewGetter(in io.Reader) (getter *Getter) {
	getter = &Getter{in: in}
	return
}

func NewStdinGetter() (getter *Getter) {
	getter = NewGetter(os.Stdin)
	return
}

func (g *Getter) Fscan() (input string) {
	fmt.Fscan(g.in, &input)
	return
}
