package iohelpers

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type noInputNoReturnFunc func()

func CaptureOutput(operation noInputNoReturnFunc) string {
	reader, writer, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = writer

	operation()

	os.Stdout = stdout
	writer.Close()
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during io.Copy: %v", err)
	}
	reader.Close()

	return buf.String()
}

func SimulateInput(stubbedInput string, operation noInputNoReturnFunc) {
	SimulateMultipleInput([]string{stubbedInput}, operation)
}

func SimulateMultipleInput(inputSequence []string, operation noInputNoReturnFunc) {
	reader, writer, _ := os.Pipe()
	byt := bytes.TrimSpace([]byte(strings.Join(inputSequence, "\n")))
	buf := bytes.NewBuffer(byt)

	_, err := io.Copy(writer, buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred during io.Copy: %v", err)
	}
	writer.Close()

	stdin := os.Stdin
	os.Stdin = reader

	operation()

	os.Stdin = stdin
	reader.Close()
}
