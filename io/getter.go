package io

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

func (g *Getter) Fscanln() (input string) {
	fmt.Fscanln(g.in, &input)
	return
}

func (g *Getter) GetString() (input string) {
	input = g.Fscanln()
	return
}

func (g *Getter) GetInt() (input int) {
	input = g.GetIntWithCallback(func() {})
	return
}

type errCallback func()

func (g *Getter) GetIntWithCallback(cb errCallback) (input int) {
	str := g.GetString()
	i64, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		cb()
		input = g.GetIntWithCallback(cb)
	} else {
		input = int(i64)
	}
	return
}
