package mocks

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
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
	missionControl := args.Get(0)
	if missionControl == nil {
		return nil, args.Error(1)
	}
	return missionControl.(*domain.MissionControl), args.Error(1)
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

func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*domain.User), args.Error(1)
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
