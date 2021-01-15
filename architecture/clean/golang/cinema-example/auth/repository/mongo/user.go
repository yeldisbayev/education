package mock

import (
	"context"

	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/entity"
)

// UserStorageMongo struct
type UserStorageMongo struct{}

// NewUserRepository constructor
func NewUserRepository() *UserStorageMongo {
	return &UserStorageMongo{}
}

// CreateUser implementation
func (s *UserStorageMongo) CreateUser(ctx context.Context, user *entity.User) error {
	return nil
}

// GetUser implementation
func (s *UserStorageMongo) GetUser(ct context.Context, username, password string) (*entity.User, error) {
	return &entity.User{
		Username: username,
		Password: password,
	}, nil
}
