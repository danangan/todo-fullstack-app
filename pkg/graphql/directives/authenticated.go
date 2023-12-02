package directives

import (
	"app/graph/generated"
	appContext "app/pkg/app-context"
	"app/pkg/db/models"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func NewDirectiveRoot() *generated.DirectiveRoot {
	return &generated.DirectiveRoot{
		Authenticated: authenticated,
	}
}

func authenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	val := ctx.Value(appContext.CurrentUserKey)

	currentUser := val.(*models.User)
	fmt.Println(currentUser)
	return nil, nil
}
