package main

import (
	"fmt"
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
	"github.com/spadin/tictactoe/player"
)

type Tictactoe struct {
	board    game.Board
	rules    game.Rules
	players  player.Players
	prompter io.Prompter
}

func NewTictactoe(prompter io.Prompter) (t *Tictactoe) {
	board := game.NewTictactoeBoard()
	rules := game.NewTictactoeRules(board)
	players := player.NewTictactoePlayers()

	t = &Tictactoe{board: board, rules: rules, players: players, prompter: prompter}
	return
}

func (t *Tictactoe) PromptForPlayers() (string, string) {
	X := t.prompter.PromptChoiceList("What is X?", t.players.AvailablePlayerTypeChoices()...)
	O := t.prompter.PromptChoiceList("What is O?", t.players.AvailablePlayerTypeChoices()...)
	return X, O
}

func (t *Tictactoe) PromptAndSetPlayers() {
	X, O := t.PromptForPlayers()
	t.players.MakeAndAdd(X, O, t.board, t.prompter)
}

func (t *Tictactoe) PrintBoard() {
	t.prompter.Printer().Print(t.board)
}

func (t *Tictactoe) PlayGame() {
	currentPlayer := t.players.Next(nil)
	for !t.rules.IsGameover() {
		t.PrintBoard()
		t.board.SetMark(currentPlayer.Mark(), currentPlayer.Play())
		currentPlayer = t.players.Next(currentPlayer)
	}
	t.PrintBoard()
}

func (t *Tictactoe) GameoverMessage() (message string) {
	if winner := t.rules.Winner(); winner == nil {
		message = "Game over, tied game"
	} else {
		message = fmt.Sprintf("Game over, %v has won the game", winner)
	}
	return
}

func (t *Tictactoe) PrintGameoverMessage() {
	message := t.GameoverMessage()
	t.prompter.Printer().Println(message)
}

func (t *Tictactoe) PromptPlayAgain() (playAgain bool) {
	command := t.prompter.PromptChoiceList("What would you like to do?", "play again", "quit")
	if command == "play again" {
		playAgain = true
	}
	return
}

func play(prompter io.Prompter) bool {
	game := NewTictactoe(prompter)
	game.PromptAndSetPlayers()
	game.PlayGame()
	game.PrintGameoverMessage()
	return game.PromptPlayAgain()
}

func main() {
	prompter := io.NewStdPrompter()

	firstTime := true
	for firstTime || play(prompter) {
		firstTime = false
	}

	prompter.Printer().Println("Goodbye!")
}
