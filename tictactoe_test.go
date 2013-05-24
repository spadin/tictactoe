package main

import (
	"github.com/remogatto/prettytest"
	"github.com/spadin/tictactoe/game"
	"github.com/spadin/tictactoe/io"
	"github.com/spadin/tictactoe/player"
	"testing"
)

type testTictactoeSuite struct {
	prettytest.Suite
}

func suiteSetup() (t *Tictactoe) {
	prompter := io.NewStdPrompter()
	t = NewTictactoe(prompter)
	return
}

func TestTictactoeRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testTictactoeSuite),
	)
}

func (t *testTictactoeSuite) TestPromptForPlayers() {
	var X, O string
	io.CaptureOutput(func() {
		io.SimulateMultipleInput([]string{"1", "1"}, func() {
			ttt := suiteSetup()
			X, O = ttt.PromptForPlayers()
		})
	})
	t.Equal("human", X)
	t.Equal("human", O)
}

func (t *testTictactoeSuite) TestPromptAndSetPlayers() {
	var ttt *Tictactoe
	io.CaptureOutput(func() {
		io.SimulateMultipleInput([]string{"1", "1"}, func() {
			ttt = suiteSetup()
			ttt.PromptAndSetPlayers()
		})
	})
	_, p1err := ttt.players.List()[0].(*player.Human)
	_, p2err := ttt.players.List()[1].(*player.Human)

	t.Nil(p1err)
	t.Nil(p2err)
}

type MockBoard struct{}

func (m *MockBoard) IsEmpty(i int) bool {
	return false
}
func (m *MockBoard) IsFull() bool {
	return false
}
func (m *MockBoard) OpenPositions() []int {
	return nil
}
func (m *MockBoard) SetMark(mark game.Mark, pos int) {
	return
}
func (m *MockBoard) String() string {
	return "mock-board"
}

func (t *testTictactoeSuite) TestPrintBoard() {
	var ttt *Tictactoe
	output := io.CaptureOutput(func() {
		io.SimulateMultipleInput([]string{"1", "1"}, func() {
			ttt = suiteSetup()
			mockBoard := &MockBoard{}
			ttt.board = mockBoard
			ttt.PrintBoard()
		})
	})
	t.Equal("mock-board", output)
}

func (t *testTictactoeSuite) TestPlayGame() {
	var ttt *Tictactoe
	io.CaptureOutput(func() {
		io.SimulateMultipleInput([]string{"0", "4", "1", "5", "2"}, func() {
			ttt = suiteSetup()
			ttt.players.MakeAndAdd("human", "human", ttt.board, ttt.prompter)
			ttt.PlayGame()
		})
	})
	t.Equal(game.X, ttt.rules.Winner())
}

func (t *testTictactoeSuite) TestPrintGameoverMessage() {
	var ttt *Tictactoe
	output := io.CaptureOutput(func() {
		ttt = suiteSetup()
		game.PlayXWinningGame(ttt.board)
		ttt.PrintGameoverMessage()
	})
	t.Equal("Game over, X has won the game\n", output)
}

func (t *testTictactoeSuite) TestPromptPlayAgain() {
	var ttt *Tictactoe
	var playAgain bool

	io.CaptureOutput(func() {
		io.SimulateMultipleInput([]string{"1"}, func() {
			ttt = suiteSetup()
			playAgain = ttt.PromptPlayAgain()
		})
	})
	t.True(playAgain)

	io.CaptureOutput(func() {
		io.SimulateMultipleInput([]string{"2"}, func() {
			ttt = suiteSetup()
			playAgain = ttt.PromptPlayAgain()
		})
	})
	t.False(playAgain)
}
