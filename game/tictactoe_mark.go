package game

type TictactoeMark int

const E, X, O = TictactoeMark(0), TictactoeMark(1), TictactoeMark(2)

func stringify(mark TictactoeMark) (str string) {
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

func (mark TictactoeMark) String() string {
	return stringify(mark)
}
