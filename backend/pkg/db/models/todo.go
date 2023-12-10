package models

import (
	"app/graphql/generated"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Todo struct {
	BaseModel
	UserID      uuid.UUID `gorm:"not null"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title       string    `gorm:"notNull" validate:"required,min=3"`
	Description string    `gorm:"notNull" validate:"required,min=3"`
	DueDate     time.Time `gorm:"notNull" validate:"required"`
}

func (t *Todo) Validate() *validator.ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(t)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return &validationErrors
	}

	return nil
}

func (t *Todo) ToGraphTodo() *generated.Todo {
	return &generated.Todo{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserID:      t.UserID.String(),
	}
}
