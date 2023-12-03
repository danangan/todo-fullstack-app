package resolvers

import (
	"app/graphql/generated"
	"app/pkg/db/models"
	"app/pkg/jwt"
	passwordPkg "app/pkg/password"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*generated.AuthPayload, error) {

	user := &models.User{}

	result := r.Db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	err := passwordPkg.ComparePassword(password, user.Password)

	if err != nil {
		return nil, errors.New("password does not match")
	}

	token, err := jwt.GenerateToken(user.ID.String())

	if err != nil {
		fmt.Printf("failed to generate token, error: %v", err)

		return nil, errors.New("failed to generate token")
	}

	response := &generated.AuthPayload{
		Token: token,
		User:  models.DBUserToGraphUser(user),
	}

	return response, nil
}
