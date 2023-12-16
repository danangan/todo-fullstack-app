package db

import (
	"app/pkg/utils"
	"os"
)

const defaultDBHost = "localhost"
const defaultDBPort = "5432"
const defaultDBUsername = "myuser"
const defaultDBPassword = "mypassword"
const defaultDBName = "mydatabase"

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func CreateDBConfig() *DBConfig {
	host := utils.EmptyOr(os.Getenv("DB_HOST"), defaultDBHost)
	port := utils.EmptyOr(os.Getenv("DB_PORT"), defaultDBPort)
	username := utils.EmptyOr(os.Getenv("DB_USERNAME"), defaultDBUsername)
	password := utils.EmptyOr(os.Getenv("DB_PASSWORD"), defaultDBPassword)
	database := utils.EmptyOr(os.Getenv("DB_NAME"), defaultDBName)

	return &DBConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}
