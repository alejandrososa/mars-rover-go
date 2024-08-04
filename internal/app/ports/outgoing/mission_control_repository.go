package outgoing

import "github.com/alejandrososa/mars-rover-go/internal/app/core/domain"

type MissionControlRepository interface {
	Save(user *domain.User, missionControl *domain.MissionControl) error
	GetByUser(user *domain.User) (*domain.MissionControl, error)
	GetByUUID(uuid string) (*domain.MissionControl, error)
	List() ([]*domain.MissionControl, error)
}
