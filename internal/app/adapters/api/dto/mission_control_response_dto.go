package dto

type MissionControlResponse struct {
	UUID string `json:"uuid"`
}

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

type GetMissionControlResponse struct {
	Message        string                 `json:"message"`
	MissionControl MissionControlResponse `json:"mission_control"`
	Platform       PlatformResponse       `json:"platform"`
	Rovers         []RoverResponse        `json:"rovers"`
}
