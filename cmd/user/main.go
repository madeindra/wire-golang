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

	// we only need to create handler without creating repository and service, they are injected by wire
	userHandler := user.InjectUser(db)

	http.Handle("/user", userHandler.FetchByUsername())
	http.ListenAndServe(":8000", nil)
}
