package mock

import (
	"errors"

	"github.com/alejandrososa/mars-rover-go/internal/app/platform"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports"
)

type MockRover struct {
	x, y      int
	direction string
	platform  *platform.Platform
}

func NewMockRover(x, y int, direction string, platform *platform.Platform) (ports.RoverControl, error) {
	if direction != "N" && direction != "E" && direction != "S" && direction != "W" {
		return nil, errors.New("invalid direction")
	}
	return &MockRover{x: x, y: y, direction: direction, platform: platform}, nil
}

func (m *MockRover) Move() {
	nextX, nextY := m.x, m.y
	switch m.direction {
	case "N":
		nextY++
	case "E":
		nextX++
	case "S":
		nextY--
	case "W":
		nextX--
	}

	if m.platform.AllowWrapAround {
		// Wrap-around logic
		if nextX >= m.platform.Width {
			nextX = 0
		} else if nextX < 0 {
			nextX = m.platform.Width - 1
		}

		if nextY >= m.platform.Height {
			nextY = 0
		} else if nextY < 0 {
			nextY = m.platform.Height - 1
		}
	} else {
		// Boundary restriction without wrap-around
		if nextX >= m.platform.Width || nextX < 0 || nextY >= m.platform.Height || nextY < 0 {
			return
		}
	}

	// Check for obstacles
	if !m.platform.IsValidPosition(nextX, nextY) {
		// If the position is invalid (e.g., occupied by an obstacle), don't move
		return
	}

	// Update the rover's position
	m.x, m.y = nextX, nextY
}

func (m *MockRover) TurnLeft() {
	switch m.direction {
	case "N":
		m.direction = "W"
	case "W":
		m.direction = "S"
	case "S":
		m.direction = "E"
	case "E":
		m.direction = "N"
	}
}

func (m *MockRover) TurnRight() {
	switch m.direction {
	case "N":
		m.direction = "E"
	case "E":
		m.direction = "S"
	case "S":
		m.direction = "W"
	case "W":
		m.direction = "N"
	}
}

func (m *MockRover) GetPosition() (int, int) {
	return m.x, m.y
}

func (m *MockRover) GetDirection() string {
	return m.direction
}

func (m *MockRover) SetObstacles(obstacles []ports.Position) {
	m.platform.SetObstacles(obstacles)
}

func (m *MockRover) ExecuteCommand(command string) error {
	switch command {
	case "M":
		m.Move()
	case "L":
		m.TurnLeft()
	case "R":
		m.TurnRight()
	default:
		return errors.New("invalid command")
	}
	return nil
}
