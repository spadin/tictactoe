package game

type TictactoeBoard [9]Mark

func NewTictactoeBoard() (board *TictactoeBoard) {
	board = new(TictactoeBoard)
	return
}

func (board TictactoeBoard) IsEmpty(index int) bool {
	return board[index] == 0
}

func (board TictactoeBoard) IsFull() bool {
	for index := 0; index < len(board); index++ {
		if board[index] == Mark(0) {
			return false
		}
	}
	return true
}

func (board TictactoeBoard) OpenPositions() (positions []int) {
	for index, val := range board {
		if val == E {
			positions = append(positions, index)
		}
	}
	return
}

func (board *TictactoeBoard) SetMark(mark Mark, index int) {
	board[index] = mark
}
