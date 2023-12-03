package resolvers

import (
	graph "app/graphql/generated"

	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) graph.ResolverRoot {
	return &Resolver{
		Db: db,
	}
}

type Resolver struct {
	Db *gorm.DB
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
