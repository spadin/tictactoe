package player

import (
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
)

type Players interface {
	Add(Player) error
	Next(Player) Player
	MakeAndAdd(string, string, game.Board, io.Prompter)
	List() []Player
	AvailablePlayerTypeChoices() []string
}
