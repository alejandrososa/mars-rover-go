package ports

type RoverControl interface {
	Move()
	TurnLeft()
	TurnRight()
	GetPosition() (int, int)
	GetDirection() string
}
