package repositories

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type InMemoryPlatformRepository struct {
	platforms map[string]*domain.Platform
}

func NewInMemoryPlatformRepository() *InMemoryPlatformRepository {
	return &InMemoryPlatformRepository{
		platforms: make(map[string]*domain.Platform),
	}
}

func (repo *InMemoryPlatformRepository) Save(platform *domain.Platform) error {
	repo.platforms[platform.GetUUID()] = platform
	return nil
}

func (repo *InMemoryPlatformRepository) GetByUUID(uuid string) (*domain.Platform, error) {
	platform, exists := repo.platforms[uuid]
	if !exists {
		return nil, errors.New("platform not found")
	}
	return platform, nil
}

func (repo *InMemoryPlatformRepository) List() ([]*domain.Platform, error) {
	var list []*domain.Platform
	for _, platform := range repo.platforms {
		list = append(list, platform)
	}
	return list, nil
}
