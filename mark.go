package tictactoe

type Mark int

const X Mark = Mark(1)
const O Mark = Mark(2)

func stringify(mark Mark) (str string) {
	switch mark {
	case X:
		str = "X"
	case O:
		str = "O"
	case 0:
		str = " "
	}
	return
}

func (mark Mark) String() string {
	return stringify(mark)
}
