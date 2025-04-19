package game

type Player interface {
	GetName() string
	GetPosition() string
	SetPosition() string
}

type Board interface {
	GetFinalPosition(currentPosition, diceValue int) int
	IsWinningPosition(position int) bool
}

type Dice interface {
	Roll() int
}

type Game interface {
	AddPlayer(player Player)
	Play() Player
	GetCurrentPlayer() Player
	MovePlayer(diceValue int) (int, int)
	IsGameOver() bool
}
