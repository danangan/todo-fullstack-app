package resolvers

import (
	graph "app/graphql/generated"
	"context"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, title string, description string, dueDate time.Time) (*graph.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, title *string, description *string, dueDate *time.Time) (*graph.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*bool, error) {
	panic("not implemented")
}
