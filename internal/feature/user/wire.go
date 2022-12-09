//go:build wireinject
// +build wireinject

package user

import (
	"database/sql"

	"github.com/google/wire"
)

// dependency injection to create repo -> service -> handler
func InjectUser(db *sql.DB) *handler {
	panic(wire.Build(UserProviderSet))
}
