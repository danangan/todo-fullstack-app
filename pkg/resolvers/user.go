package resolvers

import (
	"app/graph/generated"
	"app/pkg/db/models"
	"context"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*generated.User, error) {
	user := &models.User{}

	result := r.Db.Where("id = ?", id).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return models.DBUserToGraphUser(user), nil
}
