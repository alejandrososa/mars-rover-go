package domain

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/google/uuid"
)

type Rover struct {
	UUID      string
	Position  common.Position
	Direction common.Direction
	Command   common.Command
	Platform  *Platform
}

// NewRover creates a new rover with the specified position and direction.
// It returns a pointer to the Rover instance.
func NewRover(x, y int, direction common.Direction, platform *Platform) (*Rover, error) {
	if direction != common.North && direction != common.East && direction != common.South && direction != common.West {
		return nil, errors.New("invalid direction")
	}
	return &Rover{
		UUID:      uuid.New().String(), // Generating a new UUID
		Position:  common.Position{X: x, Y: y},
		Direction: direction,
		Platform:  platform,
	}, nil
}

func (r *Rover) Move() {
	if r.Platform == nil {
		return
	}

	nextX, nextY := r.Position.X, r.Position.Y
	switch r.Direction {
	case common.North:
		nextY++
	case common.East:
		nextX++
	case common.South:
		nextY--
	case common.West:
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
	case common.North:
		r.Direction = common.West
	case common.West:
		r.Direction = common.South
	case common.South:
		r.Direction = common.East
	case common.East:
		r.Direction = common.North
	}
}

func (r *Rover) TurnRight() {
	switch r.Direction {
	case common.North:
		r.Direction = common.East
	case common.East:
		r.Direction = common.South
	case common.South:
		r.Direction = common.West
	case common.West:
		r.Direction = common.North
	}
}

func (r *Rover) GetUUID() string { return r.UUID }

func (r *Rover) GetPosition() common.Position {
	return r.Position
}

func (r *Rover) GetDirection() common.Direction {
	return r.Direction
}

func (r *Rover) SetObstacles(obstacles []common.Position) {
	r.Platform.SetObstacles(obstacles)
}

func (m *Rover) ExecuteCommand(command common.Command) error {
	switch command {
	case common.CommandMove:
		m.Move()
	case common.CommandLeft:
		m.TurnLeft()
	case common.CommandRight:
		m.TurnRight()
	default:
		return errors.New("invalid command")
	}
	return nil
}
