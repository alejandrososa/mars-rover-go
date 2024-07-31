package entities

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

// NewRover creates a new rover with the specified position and direction.
// It returns a pointer to the Rover instance.
func NewRover(x, y int, direction Direction) *Rover {
	return &Rover{
		Position:  Position{X: x, Y: y},
		Direction: direction,
	}
}
