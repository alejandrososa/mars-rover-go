package usecases

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/dto"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/outgoing"
)

type CreateMissionControlUseCase struct {
	PlatformRepository       outgoing.PlatformRepository
	RoverRepository          outgoing.RoverRepository
	MissionControlRepository outgoing.MissionControlRepository
	UserRepository           outgoing.UserRepository
	RoverFactory             domain.RoverFactoryInterface
	UUIDGenerator            common.UUIDGenerator
}

// NewCreateMissionControlUseCase initializes the CreateMissionControlUseCase with its dependencies.
func NewCreateMissionControlUseCase(
	platformRepository outgoing.PlatformRepository,
	roverRepository outgoing.RoverRepository,
	missionControlRepository outgoing.MissionControlRepository,
	userRepository outgoing.UserRepository,
	roverFactory domain.RoverFactoryInterface,
	uuidGenerator common.UUIDGenerator,
) *CreateMissionControlUseCase {
	return &CreateMissionControlUseCase{
		PlatformRepository:       platformRepository,
		RoverRepository:          roverRepository,
		MissionControlRepository: missionControlRepository,
		UserRepository:           userRepository,
		RoverFactory:             roverFactory,
		UUIDGenerator:            uuidGenerator,
	}
}

// Execute creates a new platform, mission control, and rovers, and saves them
func (uc *CreateMissionControlUseCase) Execute(request dto.CreateMissionControlRequest) (*dto.GetMissionControlResponse, error) {
	// Retrieve or create the user
	user := domain.NewUser(request.Username)
	user.UUID = uc.UUIDGenerator.Generate()
	if err := uc.UserRepository.Save(user); err != nil {
		return nil, err
	}

	// Create a new platform
	platform := domain.NewPlatform(request.Platform.Width, request.Platform.Height, []common.Position{})
	platform.UUID = uc.UUIDGenerator.Generate()
	if err := uc.PlatformRepository.Save(platform); err != nil {
		return nil, err
	}

	// Create MissionControl
	missionControl := domain.NewMissionControl(platform)
	missionControl.UUID = uc.UUIDGenerator.Generate()

	// Create and save the rovers
	var rovers []dto.RoverResponse
	for i := 0; i < request.Rovers.Amount; i++ {
		roverControl, err := uc.RoverFactory.NewRoverControl(0, 0, common.North, platform, missionControl.Rovers)
		if err != nil {
			return nil, errors.New("unable to create rovers: " + err.Error())
		}

		if err := missionControl.AddRover(roverControl); err != nil {
			return nil, err
		}

		rover := roverControl.(*domain.Rover)
		rover.UUID = uc.UUIDGenerator.Generate()
		if err := uc.RoverRepository.Save(rover); err != nil {
			return nil, err
		}

		rovers = append(rovers, dto.RoverResponse{
			UUID:      rover.GetUUID(),
			Position:  dto.PositionResponse{X: rover.GetPosition().X, Y: rover.GetPosition().Y},
			Direction: string(rover.GetDirection()),
		})
	}

	// Save MissionControl
	if err := uc.MissionControlRepository.Save(user, missionControl); err != nil {
		return nil, err
	}

	// Create the response
	response := &dto.GetMissionControlResponse{
		Message: "Platform created successfully",
		MissionControl: dto.MissionControlResponse{
			UUID: missionControl.GetUUID(),
		},
		Platform: dto.PlatformResponse{
			Width:  platform.Width,
			Height: platform.Height,
		},
		Rovers: rovers,
	}

	return response, nil
}
