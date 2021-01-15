package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/entity"
)

// UserStorageMock struct
type UserStorageMock struct {
	mock.Mock
}

// NewUserRepository constructor
func NewUserRepository() *UserStorageMock {
	return &UserStorageMock{}
}

// CreateUser implementation
func (s *UserStorageMock) CreateUser(ctx context.Context, user *entity.User) error {
	args := s.Called(user)
	return args.Error(0)
}

// GetUser implementation
func (s *UserStorageMock) GetUser(ct context.Context, username, password string) (*entity.User, error) {
	args := s.Called(username, password)

	return args.Get(0).(*entity.User), args.Error(1)
}
