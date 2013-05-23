package game

func isEven(i int) bool {
	return (i%2 == 0)
}

func makeMoves(board *TictactoeBoard, moves ...int) {
	for i := 0; i < len(moves); i++ {
		mark := O
		if isEven(i) {
			mark = X
		}
		board.SetMark(mark, moves[i])
	}
}

func playXWinningGame(board *TictactoeBoard) {
	makeMoves(board, 0, 4, 1, 5, 2)
}

func playTiedGame(board *TictactoeBoard) {
	makeMoves(board, 0, 1, 2, 4, 3, 5, 7, 6, 8)
}
