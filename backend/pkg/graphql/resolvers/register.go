package resolvers

import (
	"app/graphql/generated"
	"app/pkg/db/models"
	"app/pkg/jwt"
	"context"
	"fmt"
)

func (r *mutationResolver) Register(ctx context.Context, firstName string, lastName string, email string, password string) (*generated.AuthPayload, error) {
	user, err := models.NewUser(
		firstName,
		lastName,
		email,
		password,
	)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	result := r.Db.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)

		return nil, result.Error
	}

	token, err := jwt.GenerateToken(user.ID.String())

	if err != nil {
		return nil, err
	}

	response := &generated.AuthPayload{
		Token: token,
		User:  models.DBUserToGraphUser(user),
	}

	return response, nil
}
