package outgoing

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type UserRepository interface {
	Save(user *domain.User) error
	GetByUUID(uuid string) (*domain.User, error)
	List() ([]*domain.User, error)
}
