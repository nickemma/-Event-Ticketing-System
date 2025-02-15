package db

import (
	"context"
	"github.com/nickemma/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomHashPassword(),
		Name:           util.RandomName(),
		Role:           util.RandomRole(),
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

	return user
}

func TestCreateUser(t *testing.T) {
	CreateUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateUser(t)

	// Ensure CreatedAt is valid (not NULL)
	require.True(t, user1.CreatedAt.Valid, "user1.CreatedAt is invalid")

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateUser(t)

	arg := UpdateUserParams{
		ID:   user1.ID,
		Name: user1.Name,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)

}

func TestDeleteUser(t *testing.T) {
	user1 := CreateUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.Equal(t, "no rows in result set", err.Error()) // The error message returned by sql.ErrNoRows
	require.Empty(t, user2)

}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
