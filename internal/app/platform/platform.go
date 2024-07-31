package platform

import "github.com/alejandrososa/mars-rover-go/internal/app/ports"

// Platform represents the surface on which the rover is moving.
type Platform struct {
	Width, Height int
	obstacles     []ports.Position
}

// NewPlatform creates a new instance of Platform with the given dimensions and obstacles.
func NewPlatform(width, height int, obstacles []ports.Position) *Platform {
	return &Platform{Width: width, Height: height, obstacles: obstacles}
}

// IsValidPosition checks if a position (x, y) is valid on the platform
func (p *Platform) IsValidPosition(x, y int) bool {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height {
		return false
	}
	for _, obstacle := range p.obstacles {
		if obstacle.X == x && obstacle.Y == y {
			return false
		}
	}
	return true
}

// SetObstacles updates the obstacles on the platform
func (p *Platform) SetObstacles(obstacles []ports.Position) {
	p.obstacles = obstacles
}
