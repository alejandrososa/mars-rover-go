package mock

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/core/entities"
)

type MockRover struct {
	Rover *entities.Rover
}

func NewMockRover(x, y int, direction entities.Direction, platform *entities.Platform) *MockRover {
	baseRover, _ := entities.NewRover(x, y, direction, platform)
	return &MockRover{
		Rover: baseRover,
	}
}

func (m *MockRover) Move() {
	m.Rover.Move()
}

func (m *MockRover) TurnLeft() {
	m.Rover.TurnLeft()
}

func (m *MockRover) TurnRight() {
	m.Rover.TurnRight()
}

func (m *MockRover) GetPosition() entities.Position {
	return m.Rover.GetPosition()
}

func (m *MockRover) GetDirection() entities.Direction {
	return m.Rover.GetDirection()
}

func (m *MockRover) SetObstacles(obstacles []entities.Position) {
	m.Rover.SetObstacles(obstacles)
}

func (m *MockRover) ExecuteCommand(command entities.Command) error {
	return m.Rover.ExecuteCommand(command)
}
