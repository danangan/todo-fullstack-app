package appContext

import (
	"app/pkg/db/models"
	"context"
	"fmt"
)

type AppContext struct {
	key string
}

var (
	CurrentUserKey = &AppContext{key: "CurrentUser"}
)

func GetCurrentUser(ctx context.Context) (*models.User, error) {
	value := ctx.Value(CurrentUserKey)

	currentUser, ok := value.(*models.User)

	if !ok {
		fmt.Println("failed to cast current user to user model")

		return nil, fmt.Errorf("failed to get current user from context")
	}

	return currentUser, nil
}
