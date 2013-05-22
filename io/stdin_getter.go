package io

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type StdinGetter struct {
	in io.Reader
}

func NewStdinGetter() (getter *StdinGetter) {
	getter = &StdinGetter{in: os.Stdin}
	return
}

func (g *StdinGetter) Fscanln() (input string) {
	fmt.Fscanln(g.in, &input)
	return
}

func (g *StdinGetter) GetString() (input string) {
	input = g.Fscanln()
	return
}

func (g *StdinGetter) GetInt() (input int) {
	input = g.GetIntWithCallback(func() {})
	return
}

type errCallback func()

func (g *StdinGetter) GetIntWithCallback(cb errCallback) (input int) {
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
