package domain_test

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissionControl_AddRoverInOccupiedPosition(t *testing.T) {
	platform := domain.NewPlatform(10, 10, []common.Position{})
	mc := domain.NewMissionControl(platform)

	firstRover, _ := domain.NewRover(0, 0, common.North, platform)
	_ = mc.AddRover(firstRover)

	secondRover, _ := domain.NewRover(0, 0, common.North, platform)
	err := mc.AddRover(secondRover)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "position already occupied")
}

func TestMissionControl_AddRover(t *testing.T) {
	platform := domain.NewPlatform(10, 10, []common.Position{})
	missionControl := domain.NewMissionControl(platform)
	rover := &domain.Rover{
		Position:  common.Position{X: 0, Y: 0},
		Direction: common.North,
		Platform:  platform,
	}

	_ = missionControl.AddRover(rover)

	assert.Equal(t, 1, len(missionControl.Rovers))
	assert.Equal(t, rover, missionControl.Rovers[0])
}

func TestMissionControl_MoveRover(t *testing.T) {
	platform := domain.NewPlatform(10, 10, []common.Position{})
	missionControl := domain.NewMissionControl(platform)
	rover := &domain.Rover{
		Position:  common.Position{X: 0, Y: 0},
		Direction: common.North,
		Platform:  platform,
	}

	_ = missionControl.AddRover(rover)
	err := missionControl.MoveRover(0)

	assert.NoError(t, err)
	assert.Equal(t, 0, missionControl.Rovers[0].GetPosition().X)
	assert.Equal(t, 1, missionControl.Rovers[0].GetPosition().Y)
}

func TestMissionControl_CommandRover(t *testing.T) {
	platform := domain.NewPlatform(10, 10, []common.Position{})
	missionControl := domain.NewMissionControl(platform)
	rover := &domain.Rover{
		Position:  common.Position{X: 0, Y: 0},
		Direction: common.North,
		Platform:  platform,
	}

	_ = missionControl.AddRover(rover)
	err := missionControl.CommandRover(0, common.CommandMove)

	assert.NoError(t, err)
	assert.Equal(t, 0, missionControl.Rovers[0].GetPosition().X)
	assert.Equal(t, 1, missionControl.Rovers[0].GetPosition().Y)
}

func TestMissionControl_CommandRover_InvalidIndex(t *testing.T) {
	platform := domain.NewPlatform(10, 10, []common.Position{})
	missionControl := domain.NewMissionControl(platform)

	err := missionControl.CommandRover(1, common.CommandMove)

	assert.Error(t, err)
}
