package models

import (
	graphModel "app/graphql/generated"

	"github.com/go-playground/validator/v10"
)

type User struct {
	BaseModel
	FirstName string `gorm:"notNull" validate:"required,alpha,min=3"`
	LastName  string `gorm:"notNull" validate:"required,alpha,min=3"`
	Email     string `gorm:"notNull;unique" validate:"required,email"`
	Password  string `gorm:"notNull" validate:"required,min=6"`
}

func (u *User) Validate() *validator.ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(u)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return &validationErrors
	}

	return nil
}

func NewUser(firstName string, lastName string, email string, password string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}

func DBUserToGraphUser(user *User) *graphModel.User {
	return &graphModel.User{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
