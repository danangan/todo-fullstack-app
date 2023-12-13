package resolvers

import (
	graph "app/graphql/generated"
	_tokenManager "app/pkg/tokenManager"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB, redisClient *redis.Client, tokenManager *_tokenManager.TokenManager) graph.ResolverRoot {
	return &Resolver{
		Db:           db,
		RedisClient:  redisClient,
		TokenManager: tokenManager,
	}
}

type Resolver struct {
	Db           *gorm.DB
	RedisClient  *redis.Client
	TokenManager *_tokenManager.TokenManager
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
