package main

import (
	"app/pkg/db"
	"app/pkg/server"

	_ "github.com/99designs/gqlgen"
)

func main() {
	db.Migrate()

	server.CreateServer()
}
