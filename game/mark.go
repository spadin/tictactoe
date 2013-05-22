package game

type Mark int

const E, X, O = Mark(0), Mark(1), Mark(2)

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
