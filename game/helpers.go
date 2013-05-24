package game

func IsEven(i int) bool {
	return (i%2 == 0)
}

func MakeMoves(board Board, moves ...int) {
	for i := 0; i < len(moves); i++ {
		mark := O
		if IsEven(i) {
			mark = X
		}
		board.SetMark(mark, moves[i])
	}
}

func PlayXWinningGame(board Board) {
	MakeMoves(board, 0, 4, 1, 5, 2)
}

func PlayTiedGame(board Board) {
	MakeMoves(board, 0, 1, 2, 4, 3, 5, 7, 6, 8)
}
