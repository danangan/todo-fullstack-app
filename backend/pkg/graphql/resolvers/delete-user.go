package resolvers

import (
	"app/graphql/generated"
	"context"
	"fmt"
)

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*generated.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}
