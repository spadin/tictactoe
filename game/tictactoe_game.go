package game

type tictactoeGame struct {
	board Board
}

func NewTictactoeGame(board Board) (game Game) {
	game = &tictactoeGame{board: board}
	return
}

func winningCombinations() [][3]int {
	return [][3]int{
		[3]int{0, 1, 2}, [3]int{4, 5, 6}, [3]int{7, 8, 9},
		[3]int{0, 3, 6}, [3]int{1, 4, 7}, [3]int{2, 5, 8},
		[3]int{0, 4, 8}, [3]int{2, 4, 6},
	}
}

func (game *tictactoeGame) IsGameover() bool {
	return game.HasWinner() || game.board.IsFull()
}

func (game *tictactoeGame) HasWinner() bool {
	board := game.board.(*TictactoeBoard)
	for _, combination := range winningCombinations() {
		if board[combination[0]] == E {
			continue
		}
		if board[combination[0]] == board[combination[1]] &&
			board[combination[0]] == board[combination[2]] {
			return true
		}
	}
	return false
}
