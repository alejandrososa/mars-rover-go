package mock_test

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/mock"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"testing"
)

func TestMockRover_Move(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{})
	rover := mock.NewMockRover(0, 0, common.North, plt)

	rover.Move()
	position := rover.GetPosition()

	if position.X != 0 || position.Y != 1 {
		t.Errorf("Expected position (0, 1), got (%d, %d)", position.X, position.Y)
	}
}

func TestMockRover_TurnLeft(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{})
	rover := mock.NewMockRover(0, 0, common.North, plt)

	rover.TurnLeft()
	direction := rover.GetDirection()

	if direction != common.West {
		t.Errorf("Expected direction West, got %s", direction)
	}
}

func TestMockRover_TurnRight(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{})
	rover := mock.NewMockRover(0, 0, common.North, plt)

	rover.TurnRight()
	direction := rover.GetDirection()

	if direction != common.East {
		t.Errorf("Expected direction East, got %s", direction)
	}
}

func TestMockRover_SetObstacles(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{})
	rover := mock.NewMockRover(0, 0, common.North, plt)

	obstacles := []common.Position{{X: 0, Y: 1}}
	rover.SetObstacles(obstacles)
	rover.Move()
	position := rover.GetPosition()

	if position.X != 0 || position.Y != 0 {
		t.Errorf("Expected position (0, 0) due to obstacle, got (%d, %d)", position.X, position.Y)
	}
}

func TestMockRover_ExecuteCommand(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{})
	rover := mock.NewMockRover(0, 0, common.North, plt)

	err := rover.ExecuteCommand("M")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	position := rover.GetPosition()
	if position.X != 0 || position.Y != 1 {
		t.Errorf("Expected position (0, 1), got (%d, %d)", position.X, position.Y)
	}

	err = rover.ExecuteCommand("L")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	direction := rover.GetDirection()
	if direction != common.West {
		t.Errorf("Expected direction West, got %s", direction)
	}

	err = rover.ExecuteCommand("Invalid")
	if err == nil {
		t.Error("Expected error for invalid command, got nil")
	}
}
