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
		Rovers:   dto.RoversConfig{Amount: 1},
	}

	// Set a fixed UUID for the test
	fixedUUID := "fixed-uuid-for-testing"
	mockUUIDGenerator.On("Generate").Return(fixedUUID).Times(4)

	user := &domain.User{UUID: fixedUUID, Username: request.Username}
	platform := &domain.Platform{
		UUID:      fixedUUID,
		Width:     request.Platform.Width,
		Height:    request.Platform.Height,
		Obstacles: []common.Position{},
	}
	rover := &domain.Rover{
		UUID:      fixedUUID,
		Position:  common.Position{X: 0, Y: 0},
		Direction: common.North,
		Platform:  platform,
	}

	mockUserRepo.On("Save", mock.MatchedBy(func(u *domain.User) bool {
		return u.Username == user.Username
	})).Return(nil)
	mockPlatformRepo.On("Save", platform).Return(nil)
	mockRoverFactory.On("NewRoverControl", 0, 0, common.North, platform, mock.Anything).Return(rover, nil).Once()
	mockRoverRepo.On("Save", rover).Return(nil)
	mockMissionControlRepo.On("Save", mock.MatchedBy(func(u *domain.User) bool {
		return u.UUID == fixedUUID && u.Username == request.Username
	}), mock.MatchedBy(func(mc *domain.MissionControl) bool {
		return mc.UUID == fixedUUID && len(mc.Rovers) == 1
	})).Return(nil)

	// Act
	response, err := useCase.Execute(request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Platform created successfully", response.Message)
	assert.Equal(t, fixedUUID, response.MissionControl.UUID)
	assert.Equal(t, request.Platform.Width, response.Platform.Width)
	assert.Equal(t, request.Platform.Height, response.Platform.Height)
	assert.Equal(t, 1, len(response.Rovers))

	mockUserRepo.AssertExpectations(t)
	mockPlatformRepo.AssertExpectations(t)
	mockRoverFactory.AssertExpectations(t)
	mockRoverRepo.AssertExpectations(t)
	mockMissionControlRepo.AssertExpectations(t)
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
		Rovers:   dto.RoversConfig{Amount: 1},
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
