package entities_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/mock"
	platform "github.com/alejandrososa/mars-rover-go/internal/app/core/entities"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports"
)

func TestIsValidPosition(t *testing.T) {
	obstacles := []ports.Position{{X: 2, Y: 2}}
	plt := platform.NewPlatform(10, 10, obstacles)

	if !plt.IsValidPosition(1, 1) {
		t.Errorf("Expected position (1, 1) to be valid")
	}

	if plt.IsValidPosition(2, 2) {
		t.Errorf("Expected position (2, 2) to be invalid due to obstacle")
	}

	if plt.IsValidPosition(10, 10) {
		t.Errorf("Expected position (10, 10) to be invalid due to boundary")
	}

	if plt.IsValidPosition(-1, 0) {
		t.Errorf("Expected position (-1, 0) to be invalid due to boundary")
	}
}

func TestAddObstacle(t *testing.T) {
	plt := platform.NewPlatform(10, 10, []ports.Position{})
	newObstacle := ports.Position{X: 5, Y: 5}
	plt.SetObstacles([]ports.Position{newObstacle})

	if plt.IsValidPosition(5, 5) {
		t.Errorf("Expected position (5, 5) to be invalid due to new obstacle")
	}
}

func TestBoundaryConditions(t *testing.T) {
	plt := platform.NewPlatform(10, 10, []ports.Position{})

	// Test edges
	if !plt.IsValidPosition(0, 0) {
		t.Errorf("Expected position (0, 0) to be valid")
	}

	if !plt.IsValidPosition(9, 9) {
		t.Errorf("Expected position (9, 9) to be valid")
	}

	// Test outside boundaries
	if plt.IsValidPosition(-1, -1) {
		t.Errorf("Expected position (-1, -1) to be invalid due to boundary")
	}

	if plt.IsValidPosition(10, 10) {
		t.Errorf("Expected position (10, 10) to be invalid due to boundary")
	}
}

func TestRemoveObstacles(t *testing.T) {
	obstacles := []ports.Position{{X: 1, Y: 1}, {X: 2, Y: 2}}
	plt := platform.NewPlatform(10, 10, obstacles)

	// Remove all obstacles
	plt.SetObstacles([]ports.Position{})

	// Ensure previous obstacle positions are now valid
	if !plt.IsValidPosition(1, 1) {
		t.Errorf("Expected position (1, 1) to be valid after removing obstacles")
	}

	if !plt.IsValidPosition(2, 2) {
		t.Errorf("Expected position (2, 2) to be valid after removing obstacles")
	}
}

func TestAllowWrapAroundTrue(t *testing.T) {
	plt := platform.NewPlatform(10, 10, []ports.Position{}, true)

	// Test wrap-around on the Y-axis
	rover, _ := mock.NewMockRover(5, 9, "N", plt)
	rover.Move()
	x, y := rover.GetPosition()

	if x != 5 || y != 0 {
		t.Errorf("Expected position to wrap-around to (5, 0), got (%d, %d)", x, y)
	}

	// Test wrap-around on the X-axis
	rover, _ = mock.NewMockRover(9, 5, "E", plt)
	rover.Move()
	x, y = rover.GetPosition()

	if x != 0 || y != 5 {
		t.Errorf("Expected position to wrap-around to (0, 5), got (%d, %d)", x, y)
	}
}

func TestAllowWrapAroundFalse(t *testing.T) {
	plt := platform.NewPlatform(10, 10, []ports.Position{}, false)

	// Test no wrap-around on the X-axis
	x, y := plt.Width-1, 5
	if plt.IsValidPosition(x+1, y) {
		t.Errorf("Expected position to not be valid beyond the boundary (10, 5)")
	}

	// Test no wrap-around on the Y-axis
	x, y = 5, plt.Height-1
	if plt.IsValidPosition(x, y+1) {
		t.Errorf("Expected position to not be valid beyond the boundary (5, 10)")
	}
}
