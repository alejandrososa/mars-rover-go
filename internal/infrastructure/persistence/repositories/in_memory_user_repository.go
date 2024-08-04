package repositories

import (
	"errors"
	"github.com/alejandrososa/mars-rover-go/internal/app/core/domain"
)

type InMemoryUserRepository struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]*domain.User)}
}

func (repo *InMemoryUserRepository) Save(user *domain.User) error {
	repo.users[user.UUID] = user
	return nil
}

func (repo *InMemoryUserRepository) GetByUUID(uuid string) (*domain.User, error) {
	user, exists := repo.users[uuid]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (repo *InMemoryUserRepository) GetByUsername(username string) (*domain.User, error) {
	for _, user := range repo.users {
		if user.GetUsername() == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepository) List() ([]*domain.User, error) {
	var users []*domain.User
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users, nil
}
