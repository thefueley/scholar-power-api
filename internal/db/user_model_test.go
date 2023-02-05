// build +integration
package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)
	var username = "test"
	var password = "test"

	usr, err := db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
}

func TestGetUserByID(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	usr, err := db.GetUserByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usr.ID, "1")
}

func TestGetUserByName(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	usr, err := db.GetUserByName(context.Background(), "test")
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usr.UserName, "test")
}

func TestUpdateUserPassword(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	passwordChangeErr := db.UpdateUserPassword(context.Background(), "test", "test")
	require.NoError(t, passwordChangeErr)
}

func TestDeleteUser(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	deleteUserErr := db.DeleteUser(context.Background(), "test")
	require.NoError(t, deleteUserErr)
}
