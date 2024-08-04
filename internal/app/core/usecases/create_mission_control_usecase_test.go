package usecases_test

import (
	"errors"
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUUIDGenerator is a mock for UUID generation.
type MockUUIDGenerator struct {
	mock.Mock
}

func (m *MockUUIDGenerator) Generate() string {
	args := m.Called()
	return args.String(0)
}

// MockPlatformRepository is a mock for the platform repository.
type MockPlatformRepository struct {
	mock.Mock
}

func (m *MockPlatformRepository) Save(platform *domain.Platform) error {
	args := m.Called(platform)
	return args.Error(0)
}

func (m *MockPlatformRepository) GetByUUID(uuid string) (*domain.Platform, error) {
	args := m.Called(uuid)
	return args.Get(0).(*domain.Platform), args.Error(1)
}

func (m *MockPlatformRepository) List() ([]*domain.Platform, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Platform), args.Error(1)
}

// MockRoverRepository is a mock for the rover repository.
type MockRoverRepository struct {
	mock.Mock
}

func (m *MockRoverRepository) Save(rover *domain.Rover) error {
	args := m.Called(rover)
	return args.Error(0)
}

func (m *MockRoverRepository) GetByUUID(uuid string) (*domain.Rover, error) {
	args := m.Called(uuid)
	return args.Get(0).(*domain.Rover), args.Error(1)
}

func (m *MockRoverRepository) List() ([]*domain.Rover, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Rover), args.Error(1)
}

// MockMissionControlRepository is a mock for the mission control repository.
type MockMissionControlRepository struct {
	mock.Mock
}

func (m *MockMissionControlRepository) Save(user *domain.User, missionControl *domain.MissionControl) error {
	args := m.Called(user, missionControl)
	return args.Error(0)
}

func (m *MockMissionControlRepository) GetByUser(user *domain.User) (*domain.MissionControl, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.MissionControl), args.Error(1)
}

func (m *MockMissionControlRepository) GetByUUID(uuid string) (*domain.MissionControl, error) {
	args := m.Called(uuid)
	return args.Get(0).(*domain.MissionControl), args.Error(1)
}

func (m *MockMissionControlRepository) List() ([]*domain.MissionControl, error) {
	args := m.Called()
	return args.Get(0).([]*domain.MissionControl), args.Error(1)
}

// MockUserRepository is a mock for the user repository.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByUUID(uuid string) (*domain.User, error) {
	args := m.Called(uuid)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) List() ([]*domain.User, error) {
	args := m.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

// MockRoverFactory is a mock for the rover factory.
type MockRoverFactory struct {
	mock.Mock
}

func (f *MockRoverFactory) NewRoverControl(x, y int, direction common.Direction, platform *domain.Platform, rovers []domain.RoverControl) (domain.RoverControl, error) {
	args := f.Called(x, y, direction, platform, rovers)
	if args.Get(0) != nil {
		return args.Get(0).(domain.RoverControl), args.Error(1)
	}
	return nil, args.Error(1)
}

// TestCreatePlatformUseCase_Execute_Success verifies that the platform and rovers are created correctly.
func TestCreatePlatformUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockPlatformRepo := new(MockPlatformRepository)
	mockRoverRepo := new(MockRoverRepository)
	mockMissionControlRepo := new(MockMissionControlRepository)
	mockUserRepo := new(MockUserRepository)
	mockRoverFactory := new(MockRoverFactory)
	mockUUIDGenerator := new(MockUUIDGenerator)

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
	mockPlatformRepo := new(MockPlatformRepository)
	mockRoverRepo := new(MockRoverRepository)
	mockMissionControlRepo := new(MockMissionControlRepository)
	mockUserRepo := new(MockUserRepository)
	mockRoverFactory := new(MockRoverFactory)
	mockUUIDGenerator := new(MockUUIDGenerator)

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
