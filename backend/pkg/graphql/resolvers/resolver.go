package resolvers

import (
	graph "app/graphql/generated"
	todoservice "app/pkg/todoService"
	"app/pkg/tokenService"
	"app/pkg/userService"
)

func NewResolver(tokenService *tokenService.TokenManager, userService *userService.UserService, todoService *todoservice.TodoService) graph.ResolverRoot {
	return &Resolver{
		TokenService: tokenService,
		UserService:  userService,
		TodoService:  todoService,
	}
}

type Resolver struct {
	TokenService *tokenService.TokenManager
	UserService  *userService.UserService
	TodoService  *todoservice.TodoService
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
