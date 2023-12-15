package userService

import (
	"app/pkg/appError"
	"app/pkg/db/models"
	"app/pkg/password"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (u *UserService) GetUser(id string) (*models.User, error) {
	user := &models.User{}

	result := u.Db.Where("id = ?", id).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserService) CreateUser(firstName string, lastName string, email string, passwordString string) (*models.User, error) {
	user := models.NewUser(
		firstName,
		lastName,
		email,
		passwordString,
	)

	validationErrors := user.Validate()

	if validationErrors != nil {
		return nil, validationErrors
	}

	hashedPassword, err := password.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	result := u.Db.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)

		return nil, result.Error
	}

	return user, nil
}

func (u *UserService) UpdateUser(id string, firstName string, lastName string, email string) (*models.User, error) {
	user, err := u.GetUser(id)

	if err != nil {
		return nil, err
	}

	if firstName != "" {
		user.FirstName = firstName
	}

	if lastName != "" {
		user.LastName = lastName
	}

	if email != "" {
		user.Email = email
	}

	validationErrors := user.Validate()

	if validationErrors != nil {
		return nil, validationErrors
	}

	result := u.Db.Save(user)

	if result.Error != nil {
		return nil, appError.ErrServer
	}

	return user, nil
}
