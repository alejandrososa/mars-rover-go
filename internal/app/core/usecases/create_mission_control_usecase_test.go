package usecases_test

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/tests/mocks"
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreatePlatformUseCase_Execute_Success verifies that the platform and rovers are created correctly.
func TestCreatePlatformUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRoverRepo := new(mocks.MockRoverRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	useCase := usecases.NewMoveRoversUseCase(mockMissionControlRepo, mockUserRepo, mockRoverRepo)

	request := dto.MoveRoversRequest{
		Username: "test_user",
		Rovers: []dto.RoverCommand{
			{UUID: "rover-uuid-1", Commands: "LMLMLMLMM"},
			{UUID: "rover-uuid-2", Commands: "MMRMMRMRRM"},
		},
	}

	user := &domain.User{UUID: "user-uuid", Username: "test_user"}
	missionControl := &domain.MissionControl{
		UUID: "mission-control-uuid",
		Platform: &domain.Platform{
			Width:  5,
			Height: 5,
		},
		Rovers: []domain.RoverControl{},
	}
	rover1 := &domain.Rover{
		UUID:      "rover-uuid-1",
		Position:  common.Position{X: 1, Y: 2},
		Direction: common.North,
	}
	rover2 := &domain.Rover{
		UUID:      "rover-uuid-2",
		Position:  common.Position{X: 3, Y: 3},
		Direction: common.East,
	}
	missionControl.Rovers = append(missionControl.Rovers, rover1, rover2)

	mockUserRepo.On("GetByUsername", "test_user").Return(user, nil)
	mockMissionControlRepo.On("GetByUser", user).Return(missionControl, nil)
	mockRoverRepo.On("GetByUUID", "rover-uuid-1").Return(rover1, nil)
	mockRoverRepo.On("GetByUUID", "rover-uuid-2").Return(rover2, nil)
	mockRoverRepo.On("Save", rover1).Return(nil)
	mockRoverRepo.On("Save", rover2).Return(nil)

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Rovers moved successfully", response.Message)
	assert.Equal(t, 2, len(response.Rovers))

	// Verificando posición y dirección final de Rover 1
	assert.Equal(t, "rover-uuid-1", response.Rovers[0].UUID)
	assert.Equal(t, dto.PositionResponse{X: 1, Y: 3}, response.Rovers[0].Position)
	assert.Equal(t, "N", response.Rovers[0].Direction)

	// Verificando posición y dirección final de Rover 2
	assert.Equal(t, "rover-uuid-2", response.Rovers[1].UUID)
	assert.Equal(t, dto.PositionResponse{X: 5, Y: 1}, response.Rovers[1].Position)
	assert.Equal(t, "E", response.Rovers[1].Direction)

	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
	mockRoverRepo.AssertExpectations(t)
}

func TestCreatePlatformUseCase_Execute_Failure(t *testing.T) {
	// Arrange
	mockPlatformRepo := new(mocks.MockPlatformRepository)
	mockRoverRepo := new(mocks.MockRoverRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoverFactory := new(mocks.MockRoverFactory)
	mockUUIDGenerator := new(mocks.MockUUIDGenerator)

	useCase := usecases.NewCreateMissionControlUseCase(
		mockPlatformRepo,
		mockRoverRepo,
		mockMissionControlRepo,
		mockUserRepo,
		mockRoverFactory,
		mockUUIDGenerator,
	)

	request := dto.CreateMissionControlRequest{
		Username: "test_user",
		Platform: dto.PlatformDimensions{Width: 10, Height: 10},
		Rovers: []dto.RoverInitialization{
			{InitialPosition: common.Position{X: 1, Y: 2}, Direction: common.North},
		},
	}

	fixedUUID := "fixed-uuid-for-testing"
	mockUUIDGenerator.On("Generate").Return(fixedUUID)

	// Setting up the expectation for an error in the Rover creation
	expectedError := errors.New("unable to create rovers")
	mockUserRepo.On("Save", mock.Anything).Return(nil)
	mockPlatformRepo.On("Save", mock.Anything).Return(nil)
	mockRoverFactory.On("NewRoverControl", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "unable to create rovers: "+expectedError.Error(), err.Error())
	mockUserRepo.AssertExpectations(t)
	mockPlatformRepo.AssertExpectations(t)
	mockRoverFactory.AssertExpectations(t)
}
