package game

type Rules interface {
	IsGameover() bool
	HasWinner() bool
	Winner() interface{}
}
