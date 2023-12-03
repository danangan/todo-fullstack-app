package models

import (
	graphModel "app/graphql/generated"
	passwordPkg "app/pkg/password"

	"github.com/go-playground/validator/v10"
)

type User struct {
	BaseModel
	FirstName string `gorm:"notNull" validate:"required,alpha,min:3"`
	LastName  string `gorm:"notNull" validate:"required,alpha,min:3"`
	Email     string `gorm:"notNull;unique" validate:"required,email"`
	Password  string `gorm:"notNull" validate:"required,min:6"`
}

func (u *User) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(u)
}

func NewUser(firstName string, lastName string, email string, password string) (*User, error) {
	hashedPassword, err := passwordPkg.HashPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}, nil
}

func DBUserToGraphUser(user *User) *graphModel.User {
	return &graphModel.User{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
