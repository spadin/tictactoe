package player

import (
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
)

type Ai struct {
	mark     game.Mark
	board    game.Board
	prompter io.Prompter
	rules    game.Rules
}

func NewAiPlayer(mark game.Mark, board game.Board, prompter io.Prompter) (player Player) {
	rules := game.NewTictactoeRules(board)
	player = &Ai{mark: mark, board: board, prompter: prompter, rules: rules}
	return
}

func (player *Ai) Mark() game.Mark {
	return player.mark
}

func (player *Ai) opponentMark() (opponentMark game.Mark) {
	opponentMark = game.X
	if player.mark == game.X {
		opponentMark = game.O
	}
	return
}

func (player *Ai) getMark(color int) (mark game.Mark) {
	mark = player.mark
	if color == -1 {
		mark = player.opponentMark()
	}
	return
}

func (player *Ai) negamax(alpha int, beta int, color int) (score int, position int) {
	if player.rules.IsGameover() {
		return color * player.calculateScore(), -1
	}
	for _, open := range player.board.OpenPositions() {
		mark := player.getMark(color)
		player.board.SetMark(mark, open)
		val, _ := player.negamax(-beta, -alpha, -color)
		val = -val
		player.board.Undo()

		if val >= beta {
			return val, open
		}
		if val > alpha {
			alpha = val
			score = alpha
			position = open
		}
	}
	return
}

func (player *Ai) calculateScore() (score int) {
	if winner := player.rules.Winner(); winner == player.Mark() {
		score = 1
	} else if winner != nil {
		score = -1
	}
	return
}

func (player *Ai) Play() (position int) {
	_, position = player.negamax(-9999, 9999, 1)
	return
}
