package domain

import "github.com/google/uuid"

type User struct {
	UUID     string
	Username string
}

// NewUser creates a new User instance
func NewUser(username string) *User {
	return &User{
		UUID:     uuid.New().String(), // Generating a new UUID
		Username: username,
	}
}

func (u *User) GetUUID() string { return u.UUID }

func (u *User) GetUsername() string { return u.Username }
