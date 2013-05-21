package tictactoe

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func isEven(i int) bool {
	return (i%2 == 0)
}

func makeMoves(board *Board, moves ...int) {
	for i := 0; i < len(moves); i++ {
		mark := X
		if isEven(i) {
			mark = O
		}
		board.SetMark(mark, moves[i])
	}
}

func playXWinningGame(board *Board) {
	makeMoves(board, 0, 4, 1, 5, 2)
}

func playTiedGame(board *Board) {
	makeMoves(board, 0, 1, 2, 4, 3, 5, 7, 6, 8)
}

type noInputNoReturnFunc func()

func captureOutput(operation noInputNoReturnFunc) string {
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

func simulateInput(stubbedInput string, operation noInputNoReturnFunc) {
	simulateMultipleInput([]string{stubbedInput}, operation)
}

func simulateMultipleInput(inputSequence []string, operation noInputNoReturnFunc) {
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
