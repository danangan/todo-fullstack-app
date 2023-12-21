package resolvers

import (
	"app/graphql/generated"
	"app/pkg/appContext"
	"app/pkg/appError"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/gqlerror"
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

	if errors.Is(err, validator.ValidationErrors{}) {
		gqlError := gqlerror.Errorf("BOOM! Headshot")

		return nil, gqlError
	}

	if err != nil {
		return nil, appError.ErrServer
	}

	return currentUser.ToGraphUser(), nil
}
