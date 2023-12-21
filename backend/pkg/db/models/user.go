package models

import (
	graphModel "app/graphql/generated"

	"github.com/go-playground/validator/v10"
)

type User struct {
	BaseModel
	FirstName string `gorm:"notNull" validate:"required,alpha,min=3" json:"firstName"`
	LastName  string `gorm:"notNull" validate:"required,alpha,min=3" json:"lastName"`
	Email     string `gorm:"notNull;unique" validate:"required,email" json:"email"`
	Password  string `gorm:"notNull" validate:"required,min=6" json:"password"`
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

func (u *User) ToGraphUser() *graphModel.User {
	return &graphModel.User{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}

func NewUser(firstName string, lastName string, email string, password string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
