package directives

import (
	"app/graphql/generated"
	"app/pkg/appContext"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

func NewDirectiveRoot() *generated.DirectiveRoot {
	return &generated.DirectiveRoot{
		Authenticated: authenticated,
	}
}

func authenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	currentUser, err := appContext.GetCurrentUser(ctx)

	if err != nil || currentUser == nil {
		return nil, errors.New("request is not authenticated")
	}

	return next(ctx)
}
