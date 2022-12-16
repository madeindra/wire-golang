package main

import (
	"database/sql"
	"net/http"

	"github.com/madeindra/wire-golang/internal/feature/user"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// without wire, this is how we create userHandler
	// userRepository := user.NewUserRepository(db) // we need to create this function in user package,
	// userService := user.NewUserService(userRepository) // we need to create this function in user package,
	// userHandler := user.NewUserHandler(userService) // we need to create this function in user package,


	// with wire, we only need to create handler without creating repository and service, they are injected by wire
	userHandler := user.InjectUser(db) // we only to export this function, no need to create functions to create repository, service, or handler 

	http.Handle("/user", userHandler.FetchByUsername())
	http.ListenAndServe(":8000", nil)
}
