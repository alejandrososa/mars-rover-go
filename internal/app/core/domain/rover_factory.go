package domain

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
)

// RoverFactoryInterface defines the interface for creating rovers.
type RoverFactoryInterface interface {
	NewRoverControl(x, y int, direction common.Direction, platform *Platform, rovers []RoverControl) (RoverControl, error)
}

// RoverFactory is a factory that creates instances of RoverControl.
type RoverFactory struct{}

// NewRoverControl creates a new instance of RoverControl. If the position is occupied, it searches for the next available position.
func (f *RoverFactory) NewRoverControl(x, y int, direction common.Direction, platform *Platform, rovers []RoverControl) (RoverControl, error) {
	position := common.Position{X: x, Y: y}

	// Search for a free position on the platform
	for platform.IsPositionOccupied(position, rovers) {
		// Move to the next position
		if position.X < platform.Width-1 {
			position.X++
		} else if position.Y < platform.Height-1 {
			position.X = 0
			position.Y++
		} else {
			// If all possible positions are checked and none are free
			return nil, errors.New("no available positions for the rover")
		}
	}

	rover, err := NewRover(position.X, position.Y, direction, platform)
	if err != nil {
		return nil, err
	}

	return rover, nil
}
