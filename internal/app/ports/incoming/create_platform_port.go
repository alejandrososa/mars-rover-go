package incoming

import "github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"

// CreateMissionControlPort defines the methods to be implemented to handle platform creation requests.
type CreateMissionControlPort interface {
	Execute(request dto.CreateMissionControlRequest) (*dto.GetMissionControlResponse, error)
}
