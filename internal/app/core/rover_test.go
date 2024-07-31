package core_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/core"
)

func TestNewRoverInitialization(t *testing.T) {
	rover := core.NewRover(0, 0, core.North)

	if rover.Position.X != 0 || rover.Position.Y != 0 {
		t.Errorf("Expected position (0, 0), got (%d, %d)", rover.Position.X, rover.Position.Y)
	}

	if rover.Direction != 0 || core.North {
		t.Errorf("Expected direction North, got %s", rover.Direction)
	}
}
