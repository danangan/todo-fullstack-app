package resolvers

import (
	"app/graphql/generated"
	"app/pkg/db/models"
	"app/pkg/jwt"
	passwordPkg "app/pkg/password"
	"context"
	"fmt"
)

func (r *mutationResolver) Register(ctx context.Context, firstName string, lastName string, email string, password string) (*generated.AuthPayload, error) {
	user := models.NewUser(
		firstName,
		lastName,
		email,
		password,
	)

	validationErrors := user.Validate()

	if validationErrors != nil {
		return nil, validationErrors
	}

	hashedPassword, err := passwordPkg.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

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
