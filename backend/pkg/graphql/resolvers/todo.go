package resolvers

import (
	graph "app/graphql/generated"
	"app/pkg/appContext"
	"app/pkg/appError"
	"context"
	"time"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*graph.Todo, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

	todos, err := r.TodoService.GetUserTodos(user)

	if err != nil {
		return nil, appError.ErrServer
	}

	graphTodos := make([]*graph.Todo, len(todos))

	for i, v := range todos {
		graphTodos[i] = v.ToGraphTodo()
	}

	return graphTodos, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, title string, description string, dueDate time.Time) (*graph.Todo, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

	todo, err := r.TodoService.CreateTodo(user, title, description, dueDate)

	if err != nil {
		return nil, appError.ErrServer
	}

	return todo.ToGraphTodo(), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, title *string, description *string, dueDate *time.Time) (*graph.Todo, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

	todo, err := r.TodoService.UpdateTodo(id, user, title, description, dueDate)

	if err != nil {
		return nil, appError.ErrServer
	}

	return todo.ToGraphTodo(), nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	user, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return false, appError.ErrServer
	}

	success, err := r.TodoService.DeleteTodo(id, user)

	if err != nil {
		return false, appError.ErrServer
	}

	return success, nil
}
