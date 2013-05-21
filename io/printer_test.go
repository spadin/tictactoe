package io

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/io/helpers"
	"os"
	"testing"
)

type testPrinterSuite struct {
	prettytest.Suite
}

func printerSuiteSetup() (printer *Printer) {
	printer = NewStdoutPrinter()
	return
}

func TestPrinterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testPrinterSuite),
	)
}

func (t *testPrinterSuite) TestPrint() {
	out := iohelpers.CaptureOutput(func() {
		printer := printerSuiteSetup()
		printer.Print("hello world")
	})

	t.Equal("hello world", out)
}

func (t *testPrinterSuite) TestPrintln() {
	out := iohelpers.CaptureOutput(func() {
		printer := printerSuiteSetup()
		printer.Println("hello world")
	})

	t.Equal("hello world\n", out)
}

func (t *testPrinterSuite) TestPrintf() {
	out := iohelpers.CaptureOutput(func() {
		printer := printerSuiteSetup()
		printer.Printf("The sky is %s", "blue")
	})

	t.Equal("The sky is blue", out)
}

func (t *testPrinterSuite) TestNewPrinter() {
	_, writer, _ := os.Pipe()
	printer := NewPrinter(writer)
	t.Equal(writer, printer.out, "Print output set to custom writer")
}

func (t *testPrinterSuite) TestNewStdoutPrinter() {
	printer := NewStdoutPrinter()
	t.Equal(os.Stdout, printer.out, "Print output set to stdout")
}
