package tictactoe

type Board [9]Mark

func (board Board) IsEmpty(index int) bool {
	return board[index] == 0
}

func (board *Board) SetMark(mark Mark, index int) {
	board[index] = mark
}
