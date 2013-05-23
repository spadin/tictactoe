package game

type Game interface {
	IsGameover() bool
	HasWinner() bool
	Winner() Mark
}
