package dto

type PositionResponse struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlatformResponse struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type RoverResponse struct {
	UUID      string           `json:"uuid"`
	Position  PositionResponse `json:"position"`
	Direction string           `json:"direction"`
}

type CreatePlatformResponse struct {
	Message  string           `json:"message"`
	Platform PlatformResponse `json:"platform"`
	Rovers   []RoverResponse  `json:"rovers"`
}
