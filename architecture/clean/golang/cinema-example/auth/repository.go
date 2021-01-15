package auth

import (
	"context"

	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/entity"
)

// UserRepository interface
type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, username, password string) (*entity.User, error)
}
