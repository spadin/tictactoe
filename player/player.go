package player

import "github.com/spadin/tictactoe/game"

type Player interface {
	Play() int
	Mark() game.Mark
}
