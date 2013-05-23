package game

type Board interface {
	IsEmpty(int) bool
	IsFull() bool
	OpenPositions() []int
}
