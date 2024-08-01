package domain

import (
	"github.com/alejandrososa/mars-rover-go/internal/app/common"
)

// Platform represents the surface on which the rover is moving.
type Platform struct {
	Width, Height int
	Obstacles     []common.Position
	// AllowWrapAround determines whether the rover wraps around the platform's boundaries.
	// If true, the platform has no physical limits, and the rover reappears on the opposite side when crossing an edge.
	// Example: From (9, 5) moving east results in (0, 5).
	// If false, the platform has physical limits, and the rover stops at the edge without wrapping around.
	// Example: From (9, 5) moving east keeps the rover at (9, 5).
	AllowWrapAround bool
}

// NewPlatform creates a new instance of Platform with the given dimensions and obstacles.
func NewPlatform(width, height int, obstacles []common.Position, options ...bool) *Platform {
	allowWrapAround := false // value by default
	if len(options) > 0 {
		allowWrapAround = options[0]
	}
	return &Platform{
		Width:           width,
		Height:          height,
		Obstacles:       obstacles,
		AllowWrapAround: allowWrapAround,
	}
}

// IsValidPosition checks if a position (x, y) is valid on the platform
func (p *Platform) IsValidPosition(x, y int) bool {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height {
		return false
	}
	for _, obstacle := range p.Obstacles {
		if obstacle.X == x && obstacle.Y == y {
			return false
		}
	}
	return true
}

// SetObstacles updates the obstacles on the platform
func (p *Platform) SetObstacles(obstacles []common.Position) {
	p.Obstacles = obstacles
}
