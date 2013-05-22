package io

import (
	"github.com/remogatto/prettytest"
	"os"
	"testing"
)

type testStdoutPrinterSuite struct {
	prettytest.Suite
}

func printerSuiteSetup() (printer *StdoutPrinter) {
	printer = NewStdoutPrinter()
	return
}

func TestStdoutPrinterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testStdoutPrinterSuite),
	)
}

func (t *testStdoutPrinterSuite) TestPrint() {
	out := captureOutput(func() {
		printer := printerSuiteSetup()
		printer.Print("hello world")
	})

	t.Equal("hello world", out)
}

func (t *testStdoutPrinterSuite) TestPrintln() {
	out := captureOutput(func() {
		printer := printerSuiteSetup()
		printer.Println("hello world")
	})

	t.Equal("hello world\n", out)
}

func (t *testStdoutPrinterSuite) TestPrintf() {
	out := captureOutput(func() {
		printer := printerSuiteSetup()
		printer.Printf("The sky is %s", "blue")
	})

	t.Equal("The sky is blue", out)
}

func (t *testStdoutPrinterSuite) TestNewStdoutPrinter() {
	printer := NewStdoutPrinter()
	t.Equal(os.Stdout, printer.out, "Printer output set to stdout")
}
