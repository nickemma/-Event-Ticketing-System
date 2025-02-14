package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:          "test@example.com",
		HashedPassword: "hashed_password@#!23",
		Name:           pgtype.Text{String: "John Doe", Valid: true}, // ✅ Correct type
		Role:           "user",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Role, user.Role)

	require.NotEmpty(t, user.ID)
	require.NotEmpty(t, user.CreatedAt)
}
