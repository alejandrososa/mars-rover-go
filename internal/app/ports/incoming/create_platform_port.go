package incoming

import "github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"

// CreateMissionControlPort defines the methods to be implemented to handle platform creation requests.
type CreateMissionControlPort interface {
	// Execute handles the request to create a new platform.
	Execute(request dto.CreateMissionControlRequest) (*dto.CreatePlatformResponse, error)
}
