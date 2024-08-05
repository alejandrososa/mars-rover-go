package usecases

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/outgoing"
)

type MoveRoversUseCase struct {
	MissionControlRepository outgoing.MissionControlRepository
	UserRepository           outgoing.UserRepository
	RoverRepository          outgoing.RoverRepository
}

func NewMoveRoversUseCase(
	missionControlRepo outgoing.MissionControlRepository,
	userRepo outgoing.UserRepository,
	roverRepo outgoing.RoverRepository,
) *MoveRoversUseCase {
	return &MoveRoversUseCase{
		MissionControlRepository: missionControlRepo,
		UserRepository:           userRepo,
		RoverRepository:          roverRepo,
	}
}

func (uc *MoveRoversUseCase) Execute(request dto.MoveRoversRequest) (*dto.GetMissionControlResponse, error) {
	// Retrieve mission control
	user, err := uc.UserRepository.GetByUsername(request.Username)
	if err != nil || user == nil {
		return nil, errors.New("mission control not found")
	}

	missionControl, err := uc.MissionControlRepository.GetByUser(user)
	if err != nil || missionControl == nil {
		return nil, errors.New("mission control not found")
	}

	// Process each rover's commands
	for _, roverCommand := range request.Rovers {
		rover, err := uc.RoverRepository.GetByUUID(roverCommand.UUID)
		if err != nil || rover == nil {
			return nil, errors.New("rover not found")
		}
		// Execute commands
		for _, cmd := range roverCommand.Commands {
			switch common.Command(cmd) {
			case common.CommandLeft:
				rover.TurnLeft()
			case common.CommandRight:
				rover.TurnRight()
			case common.CommandMove:
				rover.Move()
			default:
				return nil, errors.New("invalid command")
			}
		}

		// Update rover position in repository
		if err := uc.RoverRepository.Save(rover); err != nil {
			return nil, err
		}
	}

	// Create the response
	return &dto.GetMissionControlResponse{
		Message: "Rovers moved successfully",
		MissionControl: dto.MissionControlResponse{
			UUID: missionControl.GetUUID(),
		},
		Platform: dto.PlatformResponse{
			Width:  missionControl.Platform.Width,
			Height: missionControl.Platform.Height,
		},
		Rovers: func() []dto.RoverResponse {
			rovers := []dto.RoverResponse{}
			for _, rover := range missionControl.Rovers {
				rovers = append(rovers, dto.RoverResponse{
					UUID:      rover.GetUUID(),
					Position:  dto.PositionResponse{X: rover.GetPosition().X, Y: rover.GetPosition().Y},
					Direction: string(rover.GetDirection()),
				})
			}
			return rovers
		}(),
	}, nil
}
