//go:build integration
// +build integration

package domain

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/codercollo/equimper"
	"github.com/codercollo/equimper/config"
	"github.com/codercollo/equimper/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	conf        *config.Config
	db          *postgres.DB
	authService equimper.AuthService
	userRepo    equimper.UserRepo
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	config.LoadEnv(".env.test")

	passwordCost = bcrypt.MinCost

	conf = config.New()

	db = postgres.New(ctx, conf)

	if err := db.Drop(); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	if err := db.Trancate(ctx); err != nil {
		log.Fatal(err)
	}

	userRepo = postgres.NewUserRepo(db)

	authService = NewAuthService(userRepo)

	code := m.Run()
	db.Close()
	os.Exit(code)

}
