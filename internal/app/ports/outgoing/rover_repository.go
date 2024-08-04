package outgoing

import "github.com/alejandrososa/mars-rover-go/internal/app/core/domain"

type RoverRepository interface {
	Save(rover *domain.Rover) error
	GetByUUID(uuid string) (*domain.Rover, error)
	List() ([]*domain.Rover, error)
}
