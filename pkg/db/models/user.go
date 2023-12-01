package models

import (
	graphModel "app/graph/generated"
	passwordPkg "app/pkg/password"
)

type User struct {
	BaseModel
	FirstName string `gorm:"notNull"`
	LastName  string `gorm:"notNull"`
	Email     string `gorm:"notNull;unique"`
	Password  string `gorm:"notNull"`
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
