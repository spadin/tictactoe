package io

type Printer interface {
	Print(...interface{})
	Println(...interface{})
	Printf(string, ...interface{})
}
