package resolver

import (
	"app/graph/generated"
	"context"
	"fmt"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*generated.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
