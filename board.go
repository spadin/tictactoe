package tictactoe

type Board [9]int

func (board Board) IsEmpty(index int) bool {
	return board[index] == 0
}

func (board *Board) SetMark(mark int, index int) {
	board[index] = mark
}
