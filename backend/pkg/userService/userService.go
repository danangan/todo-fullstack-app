package userService

import (
	"app/pkg/db/models"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (t *UserService) GetUser(id string) (*models.Todo, error) {
	return nil, nil
}

func (t *UserService) CreateUser(firstName string, lastName string, email string, password string) (*models.User, error) {
	return nil, nil
}

func (t *UserService) UpdateUser(id string, firstName string, lastName string, email string) (*models.User, error) {
	return nil, nil
}
