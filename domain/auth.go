package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/codercollo/equimper"
	"golang.org/x/crypto/bcrypt"
)

var passwordCost = bcrypt.DefaultCost

type AuthService struct {
	UserRepo equimper.UserRepo
}

func NewAuthService(ur equimper.UserRepo) *AuthService {
	return &AuthService{
		UserRepo: ur,
	}
}

func (as *AuthService) Register(ctx context.Context, input equimper.RegisterInput) (equimper.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return equimper.AuthResponse{}, err
	}

	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, equimper.ErrNotFound) {
		return equimper.AuthResponse{}, equimper.ErrUsernameTaken
	}

	if _, err := as.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, equimper.ErrNotFound) {
		return equimper.AuthResponse{}, equimper.ErrEmailTaken
	}

	user := equimper.User{
		Email:    input.Email,
		Username: input.Username,
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), passwordCost)
	if err != nil {
		return equimper.AuthResponse{}, fmt.Errorf("error hashing password: %v", err)
	}

	user.Password = string(hashPassword)

	user, err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return equimper.AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	return equimper.AuthResponse{
		AccessToken: "a token",
		User:        user,
	}, nil
}

func (as *AuthService) Login(ctx context.Context, input equimper.LoginInput) (equimper.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return equimper.AuthResponse{}, err
	}

	user, err := as.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, equimper.ErrNotFound):
			return equimper.AuthResponse{}, equimper.ErrBadCredentials
		default:
			return equimper.AuthResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return equimper.AuthResponse{}, equimper.ErrBadCredentials
	}

	return equimper.AuthResponse{
		AccessToken: "a token",
		User:        user,
	}, nil

}
