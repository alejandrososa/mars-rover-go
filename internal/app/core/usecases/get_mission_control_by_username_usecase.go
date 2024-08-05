package usecases

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/outgoing"
)

// GetMissionControlByUsernameUseCase handles the logic for retrieving mission control details for a user.
type GetMissionControlByUsernameUseCase struct {
	MissionControlRepository outgoing.MissionControlRepository
	UserRepository           outgoing.UserRepository
}

// NewGetMissionControlUseCase creates a new instance of GetMissionControlByUsernameUseCase.
func NewGetMissionControlByUsernameUseCase(
	missionControlRepo outgoing.MissionControlRepository,
	userRepo outgoing.UserRepository,
) *GetMissionControlByUsernameUseCase {
	return &GetMissionControlByUsernameUseCase{
		MissionControlRepository: missionControlRepo,
		UserRepository:           userRepo,
	}
}

// Execute retrieves the mission control information for the specified user.
func (uc *GetMissionControlByUsernameUseCase) Execute(request dto.GetMissionControlRequest) (*dto.GetMissionControlResponse, error) {
	user, err := uc.UserRepository.GetByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	missionControl, err := uc.MissionControlRepository.GetByUser(user)
	if err != nil {
		return nil, err
	}

	// Create the response
	response := &dto.GetMissionControlResponse{
		Message: "Mission Control retrieved successfully",
		MissionControl: dto.MissionControlResponse{
			UUID: missionControl.GetUUID(),
		},
		Platform: dto.PlatformResponse{
			Width:  missionControl.Platform.Width,
			Height: missionControl.Platform.Height,
		},
		Rovers: func() []dto.RoverResponse {
			updatedRovers := make([]dto.RoverResponse, len(missionControl.Rovers))
			for i, rover := range missionControl.Rovers {
				updatedRovers[i] = dto.RoverResponse{
					UUID:      rover.GetUUID(),
					Position:  dto.PositionResponse{X: rover.GetPosition().X, Y: rover.GetPosition().Y},
					Direction: string(rover.GetDirection()),
				}
			}
			return updatedRovers
		}(),
	}

	return response, nil
}
