package tictactoe

import (
	"bytes"
	"github.com/remogatto/prettytest"
	"io"
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

type captureOutputInnerFunc func()

func captureOutput(operation captureOutputInnerFunc) string {
	reader, writer, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = writer

	operation()

	os.Stdout = stdout
	writer.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, reader)
	reader.Close()

	return buf.String()
}

func (t *testPrinterSuite) TestPrint() {
	out := captureOutput(func() {
		printer := printerSuiteSetup()
		printer.Print("hello world")
	})

	t.Equal("hello world", out)
}

func (t *testPrinterSuite) TestPrintln() {
	out := captureOutput(func() {
		printer := printerSuiteSetup()
		printer.Println("hello world")
	})

	t.Equal("hello world\n", out)
}

func (t *testPrinterSuite) TestPrintf() {
	out := captureOutput(func() {
		printer := printerSuiteSetup()
		printer.Printf("The sky is %s", "blue")
	})

	t.Equal("The sky is blue", out)
}

func (t *testPrinterSuite) TestNewStdoutPrinter() {
	printer := NewStdoutPrinter()
	t.Equal(os.Stdout, printer.out, "Print output set to stdout")
}

func (t *testPrinterSuite) TestNewPrinter() {
	_, writer, _ := os.Pipe()
	printer := NewPrinter(writer)
	t.Equal(writer, printer.out, "Print output set to custom writer")
}
