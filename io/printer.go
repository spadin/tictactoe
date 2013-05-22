package io

type Printer interface {
	Print(a ...interface{})
	Println(a ...interface{})
	Printf(format string, a ...interface{})
}
