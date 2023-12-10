package resolvers

import (
	"app/graphql/generated"
	appContext "app/pkg/app-context"
	"app/pkg/db/models"
	"context"
	"fmt"
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

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*generated.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

func (r *mutationResolver) UpdateCurrentUser(ctx context.Context, firstName *string, lastName *string, email *string) (*generated.User, error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current user")
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
		return nil, fmt.Errorf("failed to update user")
	}

	return models.DBUserToGraphUser(currentUser), nil
}
