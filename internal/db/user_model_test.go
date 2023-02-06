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
	var username = "test1"
	var password = "test1"

	usr, err := db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
}

func TestGetUserByID(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	usr, err := db.GetUserByID(context.Background(), "2")
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usr.ID, "2")
}

func TestGetUserByName(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	usr, err := db.GetUserByName(context.Background(), "test1")
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usr.UserName, "test1")
}

func TestUpdateUserPassword(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	passwordChangeErr := db.UpdateUserPassword(context.Background(), "test1", "password")
	require.NoError(t, passwordChangeErr)
}

func TestLogin(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	_, err = db.Login(context.Background(), "test1", "test1")
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	deleteUserErr := db.DeleteUser(context.Background(), "test1")
	require.NoError(t, deleteUserErr)
}
