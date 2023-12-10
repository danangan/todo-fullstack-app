package models

import "github.com/google/uuid"

type Todo struct {
	BaseModel
	UserID      uuid.UUID `gorm:"not null"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title       string    `gorm:"notNull" validate:"required,alpha,min=3"`
	Description string    `gorm:"notNull" validate:"required,alpha,min=3"`
	DueDate     string    `gorm:"notNull;unique" validate:"required,email"`
	Password    string    `gorm:"notNull" validate:"required,min=6"`
}
