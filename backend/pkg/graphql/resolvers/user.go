package resolvers

import (
	"app/graphql/generated"
	"app/pkg/appContext"
	"app/pkg/appError"
	"app/pkg/db/models"
	"context"
)

func (r *queryResolver) User(ctx context.Context, id string) (*generated.User, error) {
	user := &models.User{}

	result := r.Db.Where("id = ?", id).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user.ToGraphUser(), nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*generated.User, error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, err
	}

	return currentUser.ToGraphUser(), nil
}

func (r *mutationResolver) UpdateCurrentUser(ctx context.Context, firstName *string, lastName *string, email *string) (*generated.User, error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

	if firstName != nil {
		currentUser.FirstName = *firstName
	}

	if lastName != nil {
		currentUser.LastName = *lastName
	}

	if email != nil {
		currentUser.Email = *email
	}

	validationErrors := currentUser.Validate()

	if validationErrors != nil {
		return nil, validationErrors
	}

	result := r.Db.Save(currentUser)

	if result.Error != nil {
		return nil, appError.ErrServer
	}

	return currentUser.ToGraphUser(), nil
}
