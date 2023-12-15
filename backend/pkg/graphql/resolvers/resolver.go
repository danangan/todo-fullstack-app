package resolvers

import (
	graph "app/graphql/generated"
	todoservice "app/pkg/todoService"
	"app/pkg/tokenService"
	"app/pkg/userService"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB, redisClient *redis.Client, tokenService *tokenService.TokenManager, userService *userService.UserService, todoService *todoservice.TodoService) graph.ResolverRoot {
	return &Resolver{
		Db:           db,
		RedisClient:  redisClient,
		TokenService: tokenService,
		UserService:  userService,
		TodoService:  todoService,
	}
}

type Resolver struct {
	Db           *gorm.DB
	RedisClient  *redis.Client
	TokenService *tokenService.TokenManager
	UserService  *userService.UserService
	TodoService  *todoservice.TodoService
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
