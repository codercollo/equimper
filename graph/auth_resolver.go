package graph

import (
	"context"
	"errors"

	"github.com/codercollo/equimper"
)

func mapAuthResponse(a equimper.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: a.AccessToken,
		User:        mapUser(a.User),
	}
}

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	res, err := m.AuthService.Register(ctx, equimper.RegisterInput{
		Email:           input.Email,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {
		switch {
		case errors.Is(err, equimper.ErrValidation) ||
			errors.Is(err, equimper.ErrUsernameTaken) ||
			errors.Is(err, equimper.ErrEmailTaken):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}

func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	panic("implement me")
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
	panic("implement me")
}
