package tictactoe

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
