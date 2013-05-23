package player

import (
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
)

type Human struct {
	mark     game.Mark
	board    game.Board
	prompter io.Prompter
}

func NewHumanPlayer(mark game.Mark, board game.Board, prompter io.Prompter) (player Player) {
	player = &Human{mark: mark, board: board, prompter: prompter}
	return
}

func (player *Human) Mark() game.Mark {
	return player.mark
}

func (player *Human) Play() int {
	choices := player.board.OpenPositions()
	choice := player.prompter.PromptIntChoice("Choose a position", choices...)
	return choice
}
