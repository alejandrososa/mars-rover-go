package platform_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/platform"
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
