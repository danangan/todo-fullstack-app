package main

import (
	"app/pkg/db"
	"app/pkg/server"
)

func main() {
	db.Migrate()

	server.CreateServer()
}
