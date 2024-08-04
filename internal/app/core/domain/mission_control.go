package domain

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/google/uuid"
)

type MissionControl struct {
	UUID     string
	Platform *Platform
	Rovers   []RoverControl
}

// NewMissionControl creates a new MissionControl instance
func NewMissionControl(platform *Platform) *MissionControl {
	return &MissionControl{
		UUID:     uuid.New().String(), // Generating a new UUID
		Platform: platform,
		Rovers:   []RoverControl{},
	}
}

func (mc *MissionControl) GetUUID() string { return mc.UUID }

// AddRover adds a new rover to the mission control
func (mc *MissionControl) AddRover(rover RoverControl) error {
	// Check if the position is already occupied by another rover
	for _, r := range mc.Rovers {
		if r.GetPosition() == rover.GetPosition() {
			return errors.New("position already occupied")
		}
	}
	mc.Rovers = append(mc.Rovers, rover)
	return nil
}

// MoveRover moves a specified rover forward
func (mc *MissionControl) MoveRover(index int) error {
	if index >= len(mc.Rovers) {
		return errors.New("rover index out of bounds")
	}
	mc.Rovers[index].Move()
	return nil
}

// CommandRover sends a command to the specified rover
func (mc *MissionControl) CommandRover(index int, command common.Command) error {
	if index >= len(mc.Rovers) {
		return errors.New("rover index out of bounds")
	}
	return mc.Rovers[index].ExecuteCommand(command)
}
