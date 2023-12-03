package resolvers

import (
	"app/graphql/generated"
	appContext "app/pkg/app-context"
	"app/pkg/db/models"
	"context"
)

func (r *queryResolver) User(ctx context.Context, id string) (*generated.User, error) {
	user := &models.User{}

	result := r.Db.Where("id = ?", id).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return models.DBUserToGraphUser(user), nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*generated.User, error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, err
	}

	return models.DBUserToGraphUser(currentUser), nil
}
