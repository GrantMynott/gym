package main

import (
	"database/sql"

	"github.com/GrantMynott/gym/pkg/api/v2"

	"github.com/labstack/echo/v4"
)

func main() {

	connStr := "postgres://gymuser:gympass@localhost:5432/gym?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	e := echo.New()

	server := api.NewServer(e, db)

	api.RegisterHandlers(e, api.NewStrictHandler(
		server,
		[]api.StrictMiddlewareFunc{},
	))

	e.Start(":8080")

}
