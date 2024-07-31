package platform

import "github.com/alejandrososa/mars-rover-go/internal/app/ports"

type Platform struct {
	width, height int
	obstacles     []ports.Position
}

func NewPlatform(width, height int, obstacles []ports.Position) *Platform {
	return &Platform{width: width, height: height, obstacles: obstacles}
}

func (p *Platform) IsValidPosition(x, y int) bool {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return false
	}
	for _, obstacle := range p.obstacles {
		if obstacle.X == x && obstacle.Y == y {
			return false
		}
	}
	return true
}
