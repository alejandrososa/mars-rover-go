package outgoing

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type PlatformRepository interface {
	Save(platform *domain.Platform) error
	GetByUUID(uuid string) (*domain.Platform, error)
	List() ([]*domain.Platform, error)
}
