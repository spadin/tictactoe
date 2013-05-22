package game

type Board interface {
	IsEmpty(index int) bool
	IsFull() bool
	OpenPositions() []int
}
