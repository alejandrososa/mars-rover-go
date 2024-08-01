package domain

import (
	"errors"
)

type Direction string
type Command string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

const (
	CommandMove  Command = "M"
	CommandLeft  Command = "L"
	CommandRight Command = "R"
)

type Position struct {
	X int
	Y int
}

type Rover struct {
	Position  Position
	Direction Direction
	Command   Command
	Platform  *Platform
}

// NewRover creates a new rover with the specified position and direction.
// It returns a pointer to the Rover instance.
func NewRover(x, y int, direction Direction, platform *Platform) (*Rover, error) {
	if direction != North && direction != East && direction != South && direction != West {
		return nil, errors.New("invalid direction")
	}
	return &Rover{
		Position:  Position{X: x, Y: y},
		Direction: direction,
		Platform:  platform,
	}, nil
}

func (r *Rover) Move() {
	nextX, nextY := r.Position.X, r.Position.Y
	switch r.Direction {
	case North:
		nextY++
	case East:
		nextX++
	case South:
		nextY--
	case West:
		nextX--
	}

	if r.Platform.AllowWrapAround {
		// Wrap-around logic
		if nextX >= r.Platform.Width {
			nextX = 0
		} else if nextX < 0 {
			nextX = r.Platform.Width - 1
		}

		if nextY >= r.Platform.Height {
			nextY = 0
		} else if nextY < 0 {
			nextY = r.Platform.Height - 1
		}
	} else {
		// Boundary restriction without wrap-around
		if nextX >= r.Platform.Width || nextX < 0 || nextY >= r.Platform.Height || nextY < 0 {
			return
		}
	}

	// Check for obstacles
	if !r.Platform.IsValidPosition(nextX, nextY) {
		// If the position is invalid (e.g., occupied by an obstacle), don't move
		return
	}

	// Update the rover's position
	r.Position.X, r.Position.Y = nextX, nextY
}

func (r *Rover) TurnLeft() {
	switch r.Direction {
	case North:
		r.Direction = West
	case West:
		r.Direction = South
	case South:
		r.Direction = East
	case East:
		r.Direction = North
	}
}

func (r *Rover) TurnRight() {
	switch r.Direction {
	case North:
		r.Direction = East
	case East:
		r.Direction = South
	case South:
		r.Direction = West
	case West:
		r.Direction = North
	}
}

func (r *Rover) GetPosition() Position {
	return r.Position
}

func (r *Rover) GetDirection() Direction {
	return r.Direction
}

func (r *Rover) SetObstacles(obstacles []Position) {
	r.Platform.SetObstacles(obstacles)
}

func (m *Rover) ExecuteCommand(command Command) error {
	switch command {
	case CommandMove:
		m.Move()
	case CommandLeft:
		m.TurnLeft()
	case CommandRight:
		m.TurnRight()
	default:
		return errors.New("invalid command")
	}
	return nil
}
