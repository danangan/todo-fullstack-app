package resolvers

import (
	"app/graphql/generated"
	appContext "app/pkg/app-context"
	"app/pkg/db/models"
	"context"
	"fmt"
)

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

	result := r.Db.Save(currentUser)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update user")
	}

	return models.DBUserToGraphUser(currentUser), nil
}
