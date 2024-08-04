package repositories

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type InMemoryRoverRepository struct {
	rovers map[string]*domain.Rover
}

func NewInMemoryRoverRepository() *InMemoryRoverRepository {
	return &InMemoryRoverRepository{
		rovers: make(map[string]*domain.Rover),
	}
}

func (repo *InMemoryRoverRepository) Save(rover *domain.Rover) error {
	repo.rovers[rover.GetUUID()] = rover
	return nil
}

func (repo *InMemoryRoverRepository) GetByUUID(uuid string) (*domain.Rover, error) {
	rover, exists := repo.rovers[uuid]
	if !exists {
		return nil, errors.New("rover not found")
	}
	return rover, nil
}

func (repo *InMemoryRoverRepository) List() ([]*domain.Rover, error) {
	var rovers []*domain.Rover
	for _, rover := range repo.rovers {
		rovers = append(rovers, rover)
	}
	return rovers, nil
}
