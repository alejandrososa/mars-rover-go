package mock_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/mock"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports"
)

func TestMove(t *testing.T) {
	var rover ports.RoverControl = mock.NewMockRover(0, 0, "N")
	rover.Move()
	x, y := rover.GetPosition()

	if x != 0 || y != 1 {
		t.Errorf("Expected position (0, 1), got (%d, %d)", x, y)
	}
}
