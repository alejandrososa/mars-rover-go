package domain

import "github.com/alejandrososa/mars-rover-go/internal/app/common"

type RoverControl interface {
	GetUUID() string
	Move()
	TurnLeft()
	TurnRight()
	GetPosition() common.Position
	GetDirection() common.Direction
	SetObstacles(obstacles []common.Position)
	ExecuteCommand(command common.Command) error
}
