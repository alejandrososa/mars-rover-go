package outgoing

import "github.com/alejandrososa/mars-rover-go/internal/app/common"

type RoverControl interface {
	Move()
	TurnLeft()
	TurnRight()
	GetPosition() (int, int)
	GetDirection() string
	SetObstacles(obstacles []common.Position)
	ExecuteCommand(command string) error
}
