package game

type Board [9]Mark

func NewBoard() (board *Board) {
	board = new(Board)
	return
}

func (board Board) IsEmpty(index int) bool {
	return board[index] == 0
}

func (board Board) IsFull() bool {
	for index := 0; index < len(board); index++ {
		if board[index] == Mark(0) {
			return false
		}
	}
	return true
}

func (board *Board) SetMark(mark Mark, index int) {
	board[index] = mark
}
