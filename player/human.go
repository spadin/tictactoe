package player

import (
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
)

type Human struct {
	board    game.Board
	prompter io.Prompter
}

func (player *Human) Play() int {
	choices := player.board.OpenPositions()
	choice := player.prompter.PromptIntChoice("Choose a position", choices...)
	return choice
}
