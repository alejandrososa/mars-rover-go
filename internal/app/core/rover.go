package core

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

type Position struct {
	X int
	Y int
}

type Rover struct {
	Position  Position
	Direction Direction
}

func NewRover(x, y int, direction Direction) *Rover {
	return &Rover{
		Position:  Position{X: x, Y: y},
		Direction: direction,
	}
}
