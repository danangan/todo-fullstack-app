package resolvers

import (
	graph "app/graphql/generated"
	"app/pkg/appContext"
	"app/pkg/appError"
	"app/pkg/db/models"
	"context"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, title string, description string, dueDate time.Time) (*graph.Todo, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

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

	result := r.Db.Create(todo)

	if result.Error != nil {
		return nil, appError.ErrServer
	}

	return todo.ToGraphTodo(), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, title *string, description *string, dueDate *time.Time) (*graph.Todo, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

	todo := &models.Todo{}
	result := r.Db.First(todo, "ID = ? AND UserID = ", id, user.ID.String())

	if result.Error != nil {
		return nil, appError.ErrServer
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

	result = r.Db.Save(todo)

	if result.Error != nil {
		return nil, appError.ErrServer
	}

	return todo.ToGraphTodo(), nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return false, appError.ErrServer
	}

	todo := &models.Todo{}
	result := r.Db.First(todo, "ID = ? AND UserID = ", id, user.ID.String())

	if result.Error != nil {
		return false, appError.ErrServer
	}

	result = r.Db.Delete(todo)

	if result.Error != nil {
		return false, appError.ErrServer
	}

	return true, nil
}
