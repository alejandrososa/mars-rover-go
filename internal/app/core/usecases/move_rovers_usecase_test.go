package usecases_test

import (
	"errors"
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/usecases"
	"github.com/alejandrososa/mars-rover-go/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMoveRoversUseCase_Execute_Success(t *testing.T) {
	// Arrange: Set up the necessary mocks and initial state
	mockRoverRepo := new(mocks.MockRoverRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	// Initialize the use case with the mocked repositories
	useCase := usecases.NewMoveRoversUseCase(mockMissionControlRepo, mockUserRepo, mockRoverRepo)

	// Define the input request with commands for rovers
	request := dto.MoveRoversRequest{
		Username: "test_user",
		Rovers: []dto.RoverCommand{
			{UUID: "rover-uuid-1", Commands: "LMLMLMLMM"},
			{UUID: "rover-uuid-2", Commands: "MMRMMRMRRM"},
		},
	}

	// Set up initial states: user, mission control, and rovers
	user := &domain.User{UUID: "user-uuid", Username: "test_user"}
	missionControl := &domain.MissionControl{
		UUID: "mission-control-uuid",
		Platform: &domain.Platform{
			Width:  5,
			Height: 5,
		},
		Rovers: []domain.RoverControl{}, // Initial empty; will be populated later
	}

	// Initialize rover states
	rover1 := &domain.Rover{
		UUID:      "rover-uuid-1",
		Position:  common.Position{X: 1, Y: 2},
		Direction: common.North,
		Platform:  missionControl.Platform, // Assign platform for the rover
	}
	rover2 := &domain.Rover{
		UUID:      "rover-uuid-2",
		Position:  common.Position{X: 3, Y: 3},
		Direction: common.East,
		Platform:  missionControl.Platform, // Assign platform for the rover
	}

	// Add rovers to mission control
	missionControl.Rovers = append(missionControl.Rovers, rover1, rover2)

	// Define the expected interactions and return values for mocks
	mockUserRepo.On("GetByUsername", "test_user").Return(user, nil)
	mockMissionControlRepo.On("GetByUser", user).Return(missionControl, nil)
	mockRoverRepo.On("GetByUUID", "rover-uuid-1").Return(rover1, nil)
	mockRoverRepo.On("GetByUUID", "rover-uuid-2").Return(rover2, nil)
	mockRoverRepo.On("Save", rover1).Return(nil)
	mockRoverRepo.On("Save", rover2).Return(nil)

	// Act: Execute the use case with the request
	response, err := useCase.Execute(request)

	// Assert: Validate the results
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Rovers moved successfully", response.Message)
	assert.Equal(t, 2, len(response.Rovers))

	// Verify Rover 1 final position and direction
	assert.Equal(t, "rover-uuid-1", response.Rovers[0].UUID)
	assert.Equal(t, dto.PositionResponse{X: 1, Y: 3}, response.Rovers[0].Position) // Updated position
	assert.Equal(t, "N", response.Rovers[0].Direction)

	// Verify Rover 2 final position and direction
	assert.Equal(t, "rover-uuid-2", response.Rovers[1].UUID)
	assert.Equal(t, dto.PositionResponse{X: 5, Y: 1}, response.Rovers[1].Position) // Updated position
	assert.Equal(t, "E", response.Rovers[1].Direction)

	// Ensure that the mock expectations were met
	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
	mockRoverRepo.AssertExpectations(t)
}

func TestMoveRoversUseCase_Execute_RoverNotFound(t *testing.T) {
	// Arrange
	mockRoverRepo := new(mocks.MockRoverRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	useCase := usecases.NewMoveRoversUseCase(mockMissionControlRepo, mockUserRepo, mockRoverRepo)

	request := dto.MoveRoversRequest{
		Username: "test_user",
		Rovers: []dto.RoverCommand{
			{UUID: "nonexistent-rover", Commands: "M"},
		},
	}

	user := &domain.User{UUID: "user-uuid", Username: "test_user"}
	mockUserRepo.On("GetByUsername", "test_user").Return(user, nil)
	mockRoverRepo.On("GetByUUID", "nonexistent-rover").Return(nil, errors.New("rover not found"))

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "rover not found", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
	mockRoverRepo.AssertExpectations(t)
}

func TestMoveRoversUseCase_Execute_InvalidCommand(t *testing.T) {
	// Arrange
	mockRoverRepo := new(mocks.MockRoverRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	useCase := usecases.NewMoveRoversUseCase(mockMissionControlRepo, mockUserRepo, mockRoverRepo)

	request := dto.MoveRoversRequest{
		Username: "test_user",
		Rovers: []dto.RoverCommand{
			{UUID: "rover-uuid-1", Commands: "X"},
		},
	}

	user := &domain.User{UUID: "user-uuid", Username: "test_user"}
	mockUserRepo.On("GetByUsername", "test_user").Return(user, nil)
	rover := &domain.Rover{
		UUID:      "rover-uuid-1",
		Position:  common.Position{X: 1, Y: 2},
		Direction: common.North,
	}
	mockRoverRepo.On("GetByUUID", "rover-uuid-1").Return(rover, nil)

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "invalid command", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
	mockRoverRepo.AssertExpectations(t)
}
