package dto

type PlatformDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type RoversConfig struct {
	Amount int `json:"amount"`
}

type CreateMissionControlRequest struct {
	Username        string             `json:"username"`
	Platform        PlatformDimensions `json:"platform"`
	Rovers          RoversConfig       `json:"rovers"`
	AllowWrapAround bool               `json:"allow_wrap_around"`
}
