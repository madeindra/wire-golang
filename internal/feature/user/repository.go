package user

import (
	"context"
	"database/sql"

	"github.com/madeindra/wire-golang/internal/model"
)

type UserRepository interface {
	FetchByUsername(ctx context.Context, username string) (*model.User, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (*model.User, error) {
	panic("implement this")
}
