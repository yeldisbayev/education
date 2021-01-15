package usecase

import (
	"context"

	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/auth"
	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/entity"
)

// AuthUsecase struct
type AuthUsecase struct {
	userRepo auth.UserRepository
}

// NewAuthUseCase constructor
func NewAuthUseCase(userRepo auth.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
	}
}

// SignUp implementation
func (a *AuthUsecase) SignUp(ctx context.Context, username, password string) error {
	user := &entity.User{
		Username: username,
		Password: password,
	}

	return a.userRepo.CreateUser(ctx, user)
}

// SignIn implementation
func (a *AuthUsecase) SignIn(ctx context.Context, username, password string) (string, error) {
	return a.SignIn(ctx, username, password)
}
