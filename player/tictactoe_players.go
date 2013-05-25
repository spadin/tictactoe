package player

import (
	"errors"
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
)

type constructorFunc func(game.Mark, game.Board, io.Prompter) Player

type playerType struct {
	name        string
	constructor constructorFunc
}

type TictactoePlayers struct {
	list []Player
}

func NewTictactoePlayers() (players *TictactoePlayers) {
	players = &TictactoePlayers{}
	return
}

func (players *TictactoePlayers) List() []Player {
	return players.list
}

func (players *TictactoePlayers) Add(player Player) (err error) {
	if numPlayers := len(players.list); numPlayers == 2 {
		return errors.New("Can only have two players")
	}
	players.list = append(players.list, player)
	return
}

func (players *TictactoePlayers) Next(current Player) (next Player) {
	next = players.list[0]
	if current == players.list[0] {
		next = players.list[1]
	}
	return
}

func (players *TictactoePlayers) Make(name string, mark game.Mark, board game.Board, prompter io.Prompter) (player Player) {
	constructor := players.getConstructorForPlayerTypeName(name)
	player = constructor(mark, board, prompter)
	return
}

func (players *TictactoePlayers) MakeAndAdd(XName string, OName string, board game.Board, prompter io.Prompter) {
	XPlayer := players.Make(XName, game.X, board, prompter)
	players.Add(XPlayer)

	OPlayer := players.Make(OName, game.O, board, prompter)
	players.Add(OPlayer)
}

func (players *TictactoePlayers) AvailablePlayerTypeChoices() (names []string) {
	for _, pType := range players.availablePlayerTypes() {
		names = append(names, pType.name)
	}
	return
}

func (player *TictactoePlayers) availablePlayerTypes() [](*playerType) {
	return [](*playerType){
		&playerType{name: "human", constructor: NewHumanPlayer},
		&playerType{name: "ai", constructor: NewAiPlayer},
	}
}

func (players *TictactoePlayers) getConstructorForPlayerTypeName(name string) (constructor constructorFunc) {
	constructor = nil
	for _, pType := range players.availablePlayerTypes() {
		if name == pType.name {
			return pType.constructor
		}
	}
	return

}
