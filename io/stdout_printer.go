package io

import (
	"fmt"
	"io"
	"os"
)

type StdoutPrinter struct {
	out io.Writer
}

func NewStdoutPrinter() (printer *StdoutPrinter) {
	printer = &StdoutPrinter{out: os.Stdout}
	return
}

func (p *StdoutPrinter) Print(a ...interface{}) {
	fmt.Fprint(p.out, a...)
}

func (p *StdoutPrinter) Println(a ...interface{}) {
	fmt.Fprintln(p.out, a...)
}

func (p *StdoutPrinter) Printf(format string, a ...interface{}) {
	fmt.Fprintf(p.out, format, a...)
}
