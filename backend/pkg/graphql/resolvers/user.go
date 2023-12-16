package resolvers

import (
	"app/graphql/generated"
	"app/pkg/appContext"
	"app/pkg/appError"
	"context"
)

func (r *queryResolver) CurrentUser(ctx context.Context) (*generated.User, error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, err
	}

	return currentUser.ToGraphUser(), nil
}

func (r *mutationResolver) UpdateCurrentUser(ctx context.Context, firstName *string, lastName *string) (*generated.User, error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil {
		return nil, appError.ErrServer
	}

	currentUser, err = r.UserService.UpdateUser(currentUser, firstName, lastName)

	if err != nil {
		return nil, appError.ErrServer
	}

	return currentUser.ToGraphUser(), nil
}
