package tictactoe

func isEven(i int) bool {
	return (i%2 == 0)
}

func makeMoves(board *Board, moves ...int) {
	for i := 0; i < len(moves); i++ {
		mark := O
		if isEven(i) {
			mark = X
		}
		board.SetMark(mark, moves[i])
	}
}
