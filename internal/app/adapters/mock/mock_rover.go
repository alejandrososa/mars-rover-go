package mock

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type MockRover struct {
	Rover *domain.Rover
}

func NewMockRover(x, y int, direction domain.Direction, platform *domain.Platform) *MockRover {
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

func (m *MockRover) GetPosition() domain.Position {
	return m.Rover.GetPosition()
}

func (m *MockRover) GetDirection() domain.Direction {
	return m.Rover.GetDirection()
}

func (m *MockRover) SetObstacles(obstacles []domain.Position) {
	m.Rover.SetObstacles(obstacles)
}

func (m *MockRover) ExecuteCommand(command domain.Command) error {
	return m.Rover.ExecuteCommand(command)
}
