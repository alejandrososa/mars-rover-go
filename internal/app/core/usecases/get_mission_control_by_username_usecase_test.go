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

func TestGetMissionControlUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)

	useCase := usecases.NewGetMissionControlByUsernameUseCase(mockMissionControlRepo, mockUserRepo)

	request := dto.GetMissionControlRequest{
		Username: "test_user",
	}

	user := &domain.User{UUID: "fixed-uuid-for-testing", Username: "test_user"}
	platform := &domain.Platform{
		UUID:      "fixed-platform-uuid-for-testing",
		Width:     10,
		Height:    10,
		Obstacles: []common.Position{},
	}
	missionControl := &domain.MissionControl{
		UUID:     "fixed-mc-uuid-for-testing",
		Platform: platform,
		Rovers: []domain.RoverControl{
			&domain.Rover{
				UUID:      "fixed-rover-uuid-for-testing",
				Position:  common.Position{X: 0, Y: 0},
				Direction: common.North,
				Platform:  platform,
			},
		},
	}

	// Mocking repository responses
	mockUserRepo.On("GetByUsername", "test_user").Return(user, nil)
	mockMissionControlRepo.On("GetByUser", user).Return(missionControl, nil)

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Mission Control retrieved successfully", response.Message)
	assert.Equal(t, missionControl.UUID, response.MissionControl.UUID)
	assert.Equal(t, platform.Width, response.Platform.Width)
	assert.Equal(t, platform.Height, response.Platform.Height)
	assert.Equal(t, 1, len(response.Rovers))
	assert.Equal(t, "fixed-rover-uuid-for-testing", response.Rovers[0].UUID)

	// Verify that all expectations were met
	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
}

func TestGetMissionControlUseCase_Execute_UserNotFound(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)

	useCase := usecases.NewGetMissionControlByUsernameUseCase(mockMissionControlRepo, mockUserRepo)

	request := dto.GetMissionControlRequest{
		Username: "nonexistent_user",
	}

	// Mocking repository responses for a nonexistent user
	mockUserRepo.On("GetByUsername", "nonexistent_user").Return(nil, errors.New("user not found"))

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "user not found", err.Error())

	// Verify that all expectations were met
	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
}

func TestGetMissionControlUseCase_Execute_MissionControlNotFound(t *testing.T) {
	// Arrange
	mockUserRepo := new(mocks.MockUserRepository)
	mockMissionControlRepo := new(mocks.MockMissionControlRepository)

	useCase := usecases.NewGetMissionControlByUsernameUseCase(mockMissionControlRepo, mockUserRepo)

	request := dto.GetMissionControlRequest{
		Username: "test_user",
	}

	user := &domain.User{UUID: "fixed-uuid-for-testing", Username: "test_user"}

	// Mocking repository responses for a user without mission control
	mockUserRepo.On("GetByUsername", "test_user").Return(user, nil)
	mockMissionControlRepo.On("GetByUser", user).Return(nil, errors.New("mission control not found"))

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "mission control not found", err.Error())

	// Verify that all expectations were met
	mockUserRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
}
