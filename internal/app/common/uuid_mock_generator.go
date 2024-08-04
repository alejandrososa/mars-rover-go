package common

type UUIDMockGenerator struct {
	FixedUUID string
}

func (g *UUIDMockGenerator) Generate() string {
	return g.FixedUUID
}
