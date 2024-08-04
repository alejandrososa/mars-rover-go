package incoming

import "github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"

// GetMissionControlPort defines the methods to get mission control information.
type GetMissionControlPort interface {
	Execute(request dto.GetMissionControlRequest) (*dto.GetMissionControlResponse, error)
}
