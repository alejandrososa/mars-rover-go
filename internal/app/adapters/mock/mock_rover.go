package mock

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type MockRover struct {
	Rover *domain.Rover
}

func NewMockRover(x, y int, direction common.Direction, platform *domain.Platform) *MockRover {
	baseRover, _ := domain.NewRover(x, y, direction, platform)
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

func (m *MockRover) GetPosition() common.Position {
	return m.Rover.GetPosition()
}

func (m *MockRover) GetDirection() common.Direction {
	return m.Rover.GetDirection()
}

func (m *MockRover) SetObstacles(obstacles []common.Position) {
	m.Rover.SetObstacles(obstacles)
}

func (m *MockRover) ExecuteCommand(command common.Command) error {
	return m.Rover.ExecuteCommand(command)
}
