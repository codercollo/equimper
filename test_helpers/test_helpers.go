package test_helpers

import (
	"context"
	"testing"

	"github.com/codercollo/equimper/postgres"
	"github.com/stretchr/testify/require"
)

func TeardownDB(ctx context.Context, t *testing.T, db *postgres.DB) {
	t.Helper()

	err := db.Trancate(ctx)
	require.NoError(t, err)
}
