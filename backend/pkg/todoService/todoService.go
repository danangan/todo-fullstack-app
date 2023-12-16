package todoservice

import (
	"app/pkg/db/models"
	"time"

	"gorm.io/gorm"
)

type TodoService struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *TodoService {
	return &TodoService{
		Db: db,
	}
}

func (t *TodoService) GetUserTodos(user *models.User) ([]*models.Todo, error) {
	todos := make([]*models.Todo, 0)

	result := t.Db.Find(&todos, "user_id = ?", user.ID.String())

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (t *TodoService) CreateTodo(user *models.User, title string, description string, dueDate time.Time) (*models.Todo, error) {
	todo := &models.Todo{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		User:        *user,
	}

	validationError := todo.Validate()

	if validationError != nil {
		return nil, validationError
	}

	result := t.Db.Create(todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func (t *TodoService) DeleteTodo(id string, user *models.User) (bool, error) {
	todo := &models.Todo{}

	result := t.Db.First(todo, "id = ? AND user_id = ?", id, user.ID.String())

	if result.Error != nil {
		return false, result.Error
	}

	result = t.Db.Delete(todo)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (t *TodoService) UpdateTodo(id string, user *models.User, title *string, description *string, dueDate *time.Time) (*models.Todo, error) {
	todo := &models.Todo{}

	result := t.Db.Where("id = ? AND user_id = ?", id, user.ID.String()).First(todo)

	if result.Error != nil {
		return nil, result.Error
	}

	if title != nil {
		todo.Title = *title
	}

	if description != nil {
		todo.Description = *description
	}

	if dueDate != nil {
		todo.DueDate = *dueDate
	}

	validationError := todo.Validate()

	if validationError != nil {
		return nil, validationError
	}

	result = t.Db.Save(todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}
