package todoservice

import (
	"app/pkg/db/models"
	"time"

	"gorm.io/gorm"
)

type TodoService struct {
	Db *gorm.DB
}

func (t *TodoService) GetTodo(id string) (*models.Todo, error) {
	return nil, nil
}

func (t *TodoService) GetUserTodos(userId string) ([]*models.Todo, error) {
	return nil, nil
}

func (t *TodoService) CreateTodo(userId string, title string, description string, dueDate time.Time) (*models.Todo, error) {
	return nil, nil
}

func (t *TodoService) DeleteTodo(id string) (bool, error) {
	return false, nil
}

func (t *TodoService) UpdateTodo(id string, title string, description string, dueDate time.Time) (*models.Todo, error) {
	return nil, nil
}
