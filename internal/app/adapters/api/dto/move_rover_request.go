package dto

type MoveRoversRequest struct {
	Username string         `json:"username"`
	Rovers   []RoverCommand `json:"rovers"`
}

type RoverCommand struct {
	UUID     string `json:"uuid"`
	Commands string `json:"commands"`
}
