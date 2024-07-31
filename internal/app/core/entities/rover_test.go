package entities_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/core/entities"
)

func TestNewRoverInitialization(t *testing.T) {
	rover := entities.NewRover(0, 0, entities.North)

	if rover.Position.X != 0 || rover.Position.Y != 0 {
		t.Errorf("Expected position (0, 0), got (%d, %d)", rover.Position.X, rover.Position.Y)
	}

	if rover.Direction != entities.North {
		t.Errorf("Expected direction North, got %s", rover.Direction)
	}
}
