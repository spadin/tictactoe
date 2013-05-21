package tictactoe

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	out io.Writer
}

func NewStdoutPrinter() (printer *Printer) {
	printer = NewPrinter(os.Stdout)
	return
}

func NewPrinter(out io.Writer) (printer *Printer) {
	printer = &Printer{out: out}
	return
}

func (p *Printer) Print(a ...interface{}) {
	fmt.Fprint(p.out, a...)
}

func (p *Printer) Println(a ...interface{}) {
	fmt.Fprintln(p.out, a...)
}

func (p *Printer) Printf(format string, a ...interface{}) {
	fmt.Fprintf(p.out, format, a...)
}
