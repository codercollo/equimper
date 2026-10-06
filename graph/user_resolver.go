package graph

import (
	"github.com/codercollo/equimper"
)

func mapUser(u equimper.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}
