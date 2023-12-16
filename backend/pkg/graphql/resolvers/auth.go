package resolvers

import (
	"app/graphql/generated"
	"app/pkg/appError"
	"app/pkg/db/models"
	passwordPkg "app/pkg/password"
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("server error")
var ErrUserIsARegistered = errors.New("user is already registered")

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*generated.AuthPayload, error) {
	user := &models.User{
		Email: email,
	}

	user, err := r.UserService.GetUser(user)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, appError.ErrServer
	}

	err = passwordPkg.ComparePassword(password, user.Password)

	if err != nil {
		return nil, errors.New("invalid passwword")
	}

	token, err := r.TokenService.GenerateToken(ctx, user.ID.String())

	if err != nil {
		fmt.Printf("failed to generate token, error: %v", err)

		return nil, appError.ErrServer
	}

	response := &generated.AuthPayload{
		Token: token,
		User:  user.ToGraphUser(),
	}

	return response, nil
}

func (r *mutationResolver) Register(ctx context.Context, firstName string, lastName string, email string, password string) (*generated.AuthPayload, error) {
	user, err := r.UserService.CreateUser(firstName, lastName, email, password)

	if err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			return nil, validationErr
		}

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrUserIsARegistered
		}

		return nil, appError.ErrServer
	}

	token, err := r.TokenService.GenerateToken(ctx, user.ID.String())

	if err != nil {
		return nil, err
	}

	response := &generated.AuthPayload{
		Token: token,
		User:  user.ToGraphUser(),
	}

	return response, nil
}
