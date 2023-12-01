package resolver

import (
	"app/graph/generated"
	"context"
	"fmt"
)

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, firstName *string, lastName *string, email *string) (*generated.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}
