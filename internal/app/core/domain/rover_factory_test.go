package domain_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestRoverFactory_NewRoverControl(t *testing.T) {
	platform := domain.NewPlatform(10, 10, []common.Position{})
	factory := &domain.RoverFactory{}
	var rovers []domain.RoverControl

	// Initial rover creation
	roverControl, err := factory.NewRoverControl(2, 3, common.North, platform, rovers)
	rovers = append(rovers, roverControl)

	// Assertions
	assert.NoError(t, err, "Expected no error creating a rover")
	assert.NotNil(t, roverControl, "Rover should not be nil")
	assert.Equal(t, 2, roverControl.GetPosition().X, "Expected X position to be 2")
	assert.Equal(t, 3, roverControl.GetPosition().Y, "Expected Y position to be 3")
	assert.Equal(t, common.North, roverControl.GetDirection(), "Expected direction to be North")
}
