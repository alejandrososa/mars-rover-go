package domain_test

import (
	"testing"

	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	username := "testuser"
	user := domain.NewUser(username)

	assert.NotNil(t, user, "Expected user to be non-nil")
	assert.Equal(t, username, user.GetUsername(), "Expected username to match")
	assert.NotEmpty(t, user.GetUUID(), "Expected UUID to be non-empty")
}

func TestUser_GetUUID(t *testing.T) {
	user := domain.NewUser("test_user")
	uuid := user.GetUUID()

	assert.NotEmpty(t, uuid, "Expected UUID to be non-empty")
	assert.Equal(t, user.UUID, uuid, "Expected UUID to match")
}

func TestUser_GetUsername(t *testing.T) {
	username := "test_user"
	user := domain.NewUser(username)
	retrievedUsername := user.GetUsername()

	assert.Equal(t, username, retrievedUsername, "Expected username to match")
}
