package player

type Players interface {
	Add(Player) error
	Next(Player) Player
}
