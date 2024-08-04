package common

import "github.com/google/uuid"

type UUIDRandomGenerator struct{}

func (g *UUIDRandomGenerator) Generate() string {
	return uuid.New().String()
}
