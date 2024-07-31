package entities_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/core/entities"
)

func setupTestPlatform() *entities.Platform {
	return entities.NewPlatform(10, 10, []entities.Position{})
}

func setupTestPlatformWithObstacles() *entities.Platform {
	obstacles := []entities.Position{{X: 0, Y: 1}}
	return entities.NewPlatform(10, 10, obstacles)
}

func TestMove(t *testing.T) {
	plt := setupTestPlatform()
	rover, err := entities.NewRover(0, 0, entities.North, plt)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y

	if x != 0 || y != 1 {
		t.Errorf("Expected position (0, 1), got (%d, %d)", x, y)
	}
}

func TestTurnLeft(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := entities.NewRover(0, 0, entities.North, plt)
	rover.TurnLeft()
	if rover.GetDirection() != entities.West {
		t.Errorf("Expected direction W, got %s", rover.GetDirection())
	}
}

func TestTurnRight(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := entities.NewRover(0, 0, entities.North, plt)
	rover.TurnRight()
	if rover.GetDirection() != entities.East {
		t.Errorf("Expected direction E, got %s", rover.GetDirection())
	}
}

func TestMoveAndTurn(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := entities.NewRover(0, 0, entities.North, plt)
	rover.Move()
	rover.TurnRight()
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	direction := rover.GetDirection()

	if x != 1 || y != 1 || direction != entities.East {
		t.Errorf("Expected position (1, 1) and direction E, got (%d, %d) and %s", x, y, direction)
	}
}

func TestWrapAround(t *testing.T) {
	plt := entities.NewPlatform(10, 10, []entities.Position{}, true) // allowWrapAround = true
	rover, _ := entities.NewRover(0, 9, entities.North, plt)
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y

	if x != 0 || y != 0 {
		t.Errorf("Expected position (0, 0) due to wrap-around, got (%d, %d)", x, y)
	}
}

func TestNoWrapAround(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := entities.NewRover(0, 9, entities.North, plt)
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y

	if x != 0 || y != 9 {
		t.Errorf("Expected position (0, 9) due to grid boundary, got (%d, %d)", x, y)
	}
}

func TestSequentialCommands(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := entities.NewRover(1, 2, entities.North, plt)
	rover.TurnLeft()
	rover.Move()
	rover.TurnLeft()
	rover.Move()
	rover.TurnLeft()
	rover.Move()
	rover.TurnLeft()
	rover.Move()
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	direction := rover.GetDirection()

	if x != 1 || y != 3 || direction != entities.North {
		t.Errorf("Expected position (1, 3) and direction N, got (%d, %d) and %s", x, y, direction)
	}
}

func TestObstacleEncounter(t *testing.T) {
	plt := setupTestPlatformWithObstacles()
	rover, _ := entities.NewRover(0, 0, entities.North, plt)
	rover.SetObstacles([]entities.Position{{X: 0, Y: 1}})
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	if x != 0 || y != 0 {
		t.Errorf("Expected position (0, 0) due to obstacle, got (%d, %d)", x, y)
	}
}

func TestSetObstaclesIntegration(t *testing.T) {
	plt := entities.NewPlatform(10, 10, []entities.Position{{X: 1, Y: 0}}, false) // allowWrapAround = false
	rover, _ := entities.NewRover(0, 0, entities.North, plt)

	rover.SetObstacles([]entities.Position{{X: 0, Y: 1}})
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	if x != 0 || y != 0 {
		t.Errorf("Expected position (0, 0) due to obstacle, got (%d, %d)", x, y)
	}
}

func TestInvalidDirection(t *testing.T) {
	plt := setupTestPlatform()
	// Assume the constructor returns an error if the address is invalid.
	_, err := entities.NewRover(0, 0, "InvalidDirection", plt)
	if err == nil {
		t.Error("Expected error for invalid direction, got nil")
	}
}

func TestInvalidCommand(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := entities.NewRover(0, 0, entities.North, plt)
	err := rover.ExecuteCommand("InvalidCommand")
	if err == nil {
		t.Error("Expected error for invalid command, got nil")
	}
}
