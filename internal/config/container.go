package config

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/http/controllers"
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/usecases"
	"github.com/alejandrososa/mars-rover-go/internal/app/ports/outgoing"
	"github.com/alejandrososa/mars-rover-go/internal/infrastructure/persistence/repositories"
)

// Dependency container
var (
	HealthCheckController              *controllers.HealthCheckController
	CreateMissionControlController     *controllers.CreateMissionControlController
	GetMissionControlController        *controllers.GetMissionControlController
	PlatformRepository                 outgoing.PlatformRepository
	RoverRepository                    outgoing.RoverRepository
	MissionControlRepository           outgoing.MissionControlRepository
	UserRepository                     outgoing.UserRepository
	RoverFactory                       domain.RoverFactoryInterface
	CreateMissionControlUseCase        *usecases.CreateMissionControlUseCase
	GetMissionControlByUsernameUseCase *usecases.GetMissionControlByUsernameUseCase
	UUIDGenerator                      common.UUIDGenerator
)

func init() {
	// Initialize UUID generator
	UUIDGenerator = &common.UUIDRandomGenerator{}

	// Initialize repositories
	PlatformRepository = repositories.NewInMemoryPlatformRepository()
	RoverRepository = repositories.NewInMemoryRoverRepository()
	MissionControlRepository = repositories.NewInMemoryMissionControlRepository()
	UserRepository = repositories.NewInMemoryUserRepository()

	// Initialize factories
	RoverFactory = &domain.RoverFactory{}

	// Initialize use cases
	CreateMissionControlUseCase = usecases.NewCreateMissionControlUseCase(
		PlatformRepository,
		RoverRepository,
		MissionControlRepository,
		UserRepository,
		RoverFactory,
		UUIDGenerator,
	)

	GetMissionControlByUsernameUseCase = usecases.NewGetMissionControlByUsernameUseCase(
		MissionControlRepository,
		UserRepository,
	)

	// Initialize controllers
	HealthCheckController = controllers.NewHealthCheckController()
	CreateMissionControlController = controllers.NewCreateMissionControlController(CreateMissionControlUseCase)
	GetMissionControlController = controllers.NewGetMissionControlController(GetMissionControlByUsernameUseCase)

}
