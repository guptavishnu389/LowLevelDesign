package player

import "snake_and_ladder/game"

type ConcretePlayer struct {
	name     string
	position int
}

func NewPlayer(name string) game.Player {
	return &ConcretePlayer{
		name:     name,
		position: 0,
	}
}

func (c ConcretePlayer) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (c ConcretePlayer) GetPosition() string {
	//TODO implement me
	panic("implement me")
}

func (c ConcretePlayer) SetPosition() string {
	//TODO implement me
	panic("implement me")
}
