package user

import (
	"context"

	"github.com/madeindra/wire-golang/internal/model"
)

type UserService interface {
	FetchByUsername(ctx context.Context, username string) (*model.UserResponse, error)
}

type service struct {
	r UserRepository
}

func (s *service) FetchByUsername(ctx context.Context, username string) (*model.UserResponse, error) {
	panic("implement this")
}
