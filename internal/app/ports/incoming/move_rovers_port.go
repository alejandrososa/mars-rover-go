package incoming

import "github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"

// MoveRoversPort defines the port for moving rovers.
type MoveRoversPort interface {
	Execute(request dto.MoveRoversRequest) (*dto.GetMissionControlResponse, error)
}
