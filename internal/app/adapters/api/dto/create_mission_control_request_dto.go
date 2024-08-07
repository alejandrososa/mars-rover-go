package dto

import "github.com/alejandrososa/mars-rover-go/internal/app/common"

type PlatformDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type RoverInitialization struct {
	InitialPosition common.Position  `json:"initial_position,omitempty"`
	Direction       common.Direction `json:"direction"`
}

type CreateMissionControlRequest struct {
	Username        string                `json:"username"`
	Platform        PlatformDimensions    `json:"platform"`
	Rovers          []RoverInitialization `json:"rovers"`
	AllowWrapAround bool                  `json:"allow_wrap_around"`
}
