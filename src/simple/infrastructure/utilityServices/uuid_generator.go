package utilityServices

import (
	"simple/applications/deps/services"

	"github.com/google/uuid"
)

type uuidGenerator struct{}

func NewUUIDGenerator() services.IDGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) GenerateID() string {
	return uuid.New().String()
}
