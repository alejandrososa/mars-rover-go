package repositories

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type InMemoryMissionControlRepository struct {
	data map[string]*domain.MissionControl
}

func NewInMemoryMissionControlRepository() *InMemoryMissionControlRepository {
	return &InMemoryMissionControlRepository{data: make(map[string]*domain.MissionControl)}
}

func (repo *InMemoryMissionControlRepository) Save(user *domain.User, missionControl *domain.MissionControl) error {
	repo.data[user.GetUUID()] = missionControl
	return nil
}

func (repo *InMemoryMissionControlRepository) GetByUser(user *domain.User) (*domain.MissionControl, error) {
	missionControl, exists := repo.data[user.GetUUID()]
	if !exists {
		return nil, errors.New("mission control not found")
	}
	return missionControl, nil
}

func (repo *InMemoryMissionControlRepository) GetByUUID(uuid string) (*domain.MissionControl, error) {
	missionControl, exists := repo.data[uuid]
	if !exists {
		return nil, errors.New("mission control not found")
	}
	return missionControl, nil
}

func (repo *InMemoryMissionControlRepository) List() ([]*domain.MissionControl, error) {
	var controls []*domain.MissionControl
	for _, missionControl := range repo.data {
		controls = append(controls, missionControl)
	}
	return controls, nil
}
