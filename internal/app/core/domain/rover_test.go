package domain_test

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"testing"
)

func setupTestPlatform() *domain.Platform {
	return domain.NewPlatform(10, 10, []common.Position{})
}

func setupTestPlatformWithObstacles() *domain.Platform {
	obstacles := []common.Position{{X: 0, Y: 1}}
	return domain.NewPlatform(10, 10, obstacles)
}

func TestMove(t *testing.T) {
	plt := setupTestPlatform()
	rover, err := domain.NewRover(0, 0, common.North, plt)
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
	rover, _ := domain.NewRover(0, 0, common.North, plt)
	rover.TurnLeft()
	if rover.GetDirection() != common.West {
		t.Errorf("Expected direction W, got %s", rover.GetDirection())
	}
}

func TestTurnRight(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := domain.NewRover(0, 0, common.North, plt)
	rover.TurnRight()
	if rover.GetDirection() != common.East {
		t.Errorf("Expected direction E, got %s", rover.GetDirection())
	}
}

func TestMoveAndTurn(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := domain.NewRover(0, 0, common.North, plt)
	rover.Move()
	rover.TurnRight()
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	direction := rover.GetDirection()

	if x != 1 || y != 1 || direction != common.East {
		t.Errorf("Expected position (1, 1) and direction E, got (%d, %d) and %s", x, y, direction)
	}
}

func TestWrapAround(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{}, true) // allowWrapAround = true
	rover, _ := domain.NewRover(0, 9, common.North, plt)
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y

	if x != 0 || y != 0 {
		t.Errorf("Expected position (0, 0) due to wrap-around, got (%d, %d)", x, y)
	}
}

func TestNoWrapAround(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := domain.NewRover(0, 9, common.North, plt)
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y

	if x != 0 || y != 9 {
		t.Errorf("Expected position (0, 9) due to grid boundary, got (%d, %d)", x, y)
	}
}

func TestSequentialCommands(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := domain.NewRover(1, 2, common.North, plt)
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

	if x != 1 || y != 3 || direction != common.North {
		t.Errorf("Expected position (1, 3) and direction N, got (%d, %d) and %s", x, y, direction)
	}
}

func TestMoveRoversWithCommands(t *testing.T) {
	plt := setupTestPlatform()
	rover1, _ := domain.NewRover(1, 2, common.North, plt)
	rover2, _ := domain.NewRover(3, 3, common.East, plt)

	// Commands for the first rover
	commands1 := "LMLMLMLMM"
	for _, cmd := range commands1 {
		rover1.ExecuteCommand(common.Command(cmd))
	}

	// Commands for the second rover
	commands2 := "MMRMMRMRRM"
	for _, cmd := range commands2 {
		rover2.ExecuteCommand(common.Command(cmd))
	}

	// Verify the expected outcome for the first rover
	x1, y1 := rover1.GetPosition().X, rover1.GetPosition().Y
	direction1 := rover1.GetDirection()
	if x1 != 1 || y1 != 3 || direction1 != common.North {
		t.Errorf("Expected rover1 position (1, 3) and direction N, got (%d, %d) and %s", x1, y1, direction1)
	}

	// Verify the expected outcome for the second rover
	x2, y2 := rover2.GetPosition().X, rover2.GetPosition().Y
	direction2 := rover2.GetDirection()
	if x2 != 5 || y2 != 1 || direction2 != common.East {
		t.Errorf("Expected rover2 position (5, 1) and direction E, got (%d, %d) and %s", x2, y2, direction2)
	}
}

func TestObstacleEncounter(t *testing.T) {
	plt := setupTestPlatformWithObstacles()
	rover, _ := domain.NewRover(0, 0, common.North, plt)
	rover.SetObstacles([]common.Position{{X: 0, Y: 1}})
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	if x != 0 || y != 0 {
		t.Errorf("Expected position (0, 0) due to obstacle, got (%d, %d)", x, y)
	}
}

func TestSetObstaclesIntegration(t *testing.T) {
	plt := domain.NewPlatform(10, 10, []common.Position{{X: 1, Y: 0}}, false) // allowWrapAround = false
	rover, _ := domain.NewRover(0, 0, common.North, plt)

	rover.SetObstacles([]common.Position{{X: 0, Y: 1}})
	rover.Move()
	x, y := rover.GetPosition().X, rover.GetPosition().Y
	if x != 0 || y != 0 {
		t.Errorf("Expected position (0, 0) due to obstacle, got (%d, %d)", x, y)
	}
}

func TestInvalidDirection(t *testing.T) {
	plt := setupTestPlatform()
	// Assume the constructor returns an error if the address is invalid.
	_, err := domain.NewRover(0, 0, "InvalidDirection", plt)
	if err == nil {
		t.Error("Expected error for invalid direction, got nil")
	}
}

func TestInvalidCommand(t *testing.T) {
	plt := setupTestPlatform()
	rover, _ := domain.NewRover(0, 0, common.North, plt)
	err := rover.ExecuteCommand("InvalidCommand")
	if err == nil {
		t.Error("Expected error for invalid command, got nil")
	}
}
