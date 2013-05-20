package tictactoe

type Mark int

const E Mark = Mark(0)
const X Mark = Mark(1)
const O Mark = Mark(2)

func stringify(mark Mark) (str string) {
	switch mark {
	case X:
		str = "X"
	case O:
		str = "O"
	case E:
		str = " "
	}
	return
}

func (mark Mark) String() string {
	return stringify(mark)
}
