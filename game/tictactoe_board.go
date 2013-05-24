package game

import "fmt"

type TictactoeBoard struct {
	state [9]Mark
}

func NewTictactoeBoard() (board *TictactoeBoard) {
	board = &TictactoeBoard{state: [9]Mark{E, E, E, E, E, E, E, E, E}}
	return
}

func (board TictactoeBoard) IsEmpty(index int) bool {
	return board.state[index] == E
}

func (board TictactoeBoard) IsFull() bool {
	for index := 0; index < len(board.state); index++ {
		if board.state[index] == E {
			return false
		}
	}
	return true
}

func (board TictactoeBoard) OpenPositions() (positions []int) {
	for index, val := range board.state {
		if val == E {
			positions = append(positions, index)
		}
	}
	return
}

func (board *TictactoeBoard) SetMark(mark Mark, index int) {
	board.state[index] = mark
}

func (board *TictactoeBoard) String() (str string) {
	vertBars := "     |     |     \n"
	horzBars := "-----------------\n"

	var items []interface{}
	for index, val := range board.state {
		var item interface{} = val
		if val == E {
			item = index
		}
		items = append(items, item)
	}

	str = fmt.Sprintf("\n"+vertBars+
		"  %v  |  %v  |  %v  \n"+
		vertBars+horzBars+vertBars+
		"  %v  |  %v  |  %v  \n"+
		vertBars+horzBars+vertBars+
		"  %v  |  %v  |  %v  \n"+
		vertBars, items...)
	return
}
