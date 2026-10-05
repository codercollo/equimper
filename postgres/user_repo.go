package postgres

import (
	"context"
	"fmt"

	"github.com/codercollo/equimper"
	"github.com/georgysavva/scany/pgxscan"
)

type UserRepo struct {
	DB *DB
}

func (ur *UserRepo) Create(ctx context.Context, user equimper.User) (equimper.User, error) {
	return equimper.User{}, nil
}

func (ur *UserRepo) GetByUsername(ctx context.Context, username string) (equimper.User, error) {
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1;`

	u := equimper.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, username); err != nil {
		if pgxscan.NotFound(err) {
			return equimper.User{}, equimper.ErrNotFound
		}

		return equimper.User{}, fmt.Errorf("error select: %v", err)
	}

	return u, nil
}

func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (equimper.User, error) {

	query := `SELECT * FROM users WHERE email = $1 LIMIT 1;`

	u := equimper.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, email); err != nil {
		if pgxscan.NotFound(err) {
			return equimper.User{}, equimper.ErrNotFound
		}

		return equimper.User{}, fmt.Errorf("error select: %v", err)
	}

	return u, nil
}
