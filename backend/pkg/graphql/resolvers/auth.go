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

// Todos
// * Track the user and the current active token
// * To revoke the previously granted token by adding it to the blacklist if the user still have an active token
// * Maybe we don't need to revoke it at all? Maybe we can just return the existing one?
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

	token, err := jwt.GenerateToken(ctx, r.RedisClient, user.ID.String())

	if err != nil {
		fmt.Printf("failed to generate token, error: %v", err)

		return nil, errors.New("failed to generate token")
	}

	response := &generated.AuthPayload{
		Token: token,
		User:  user.ToGraphUser(),
	}

	return response, nil
}

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

	token, err := jwt.GenerateToken(ctx, r.RedisClient, user.ID.String())

	if err != nil {
		return nil, err
	}

	response := &generated.AuthPayload{
		Token: token,
		User:  user.ToGraphUser(),
	}

	return response, nil
}
