package ports

type Position struct {
	X int
	Y int
}

type RoverControl interface {
	Move()
	TurnLeft()
	TurnRight()
	GetPosition() (int, int)
	GetDirection() string
	SetObstacles(obstacles []Position)
	ExecuteCommand(command string) error
}
