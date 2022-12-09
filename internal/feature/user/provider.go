package user

import (
	"database/sql"

	"github.com/google/wire"
)

func ProvideUserRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func ProvideUserService(repo UserRepository) *service {
	return &service{
		r: repo,
	}
}

func ProvideUserHandler(service UserService) *handler {
	return &handler{
		s: service,
	}
}

var UserProviderSet wire.ProviderSet = wire.NewSet(
	ProvideUserRepository,
	ProvideUserService,
	ProvideUserHandler,

	// bind to the interface because the provider return concrete, not interface as needed by the dependants
	wire.Bind(new(UserRepository), new(*repository)),
	wire.Bind(new(UserService), new(*service)),
	wire.Bind(new(UserHandler), new(*handler)),
)
