package models

import "github.com/google/uuid"

type Todo struct {
	BaseModel
	UserId      uuid.UUID
	User        User
	Title       string `gorm:"notNull" validate:"required,alpha,min=3"`
	Description string `gorm:"notNull" validate:"required,alpha,min=3"`
	DueDate     string `gorm:"notNull;unique" validate:"required,email"`
	Password    string `gorm:"notNull" validate:"required,min=6"`
}
