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
	var username = "test111"
	var password = "test111"

	usr, err := db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)
	require.NotEmpty(t, usr)

	usrInfo, err := db.GetUserByName(context.Background(), "test111")
	require.NoError(t, err)

	err = db.DeleteUser(context.Background(), usrInfo.ID)
	require.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	var username = "test111"
	var password = "test111"

	usr, err := db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)

	usrInfo, err := db.GetUserByName(context.Background(), "test111")
	require.NoError(t, err)

	usrI, err := db.GetUserByID(context.Background(), usrInfo.ID)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usrI.ID, usrInfo.ID)

	err = db.DeleteUser(context.Background(), usrInfo.ID)
	require.NoError(t, err)
}

func TestGetUserByName(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	var username = "test111"
	var password = "test111"

	_, err = db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)

	usr, err := db.GetUserByName(context.Background(), "test111")
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usr.UserName, "test111")

	err = db.DeleteUser(context.Background(), usr.ID)
	require.NoError(t, err)
}

func TestUpdateUserPassword(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	var username = "test111"
	var password = "test111"

	_, err = db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)

	usr, err := db.GetUserByName(context.Background(), "test111")
	require.NoError(t, err)

	passwordChangeErr := db.UpdateUserPassword(context.Background(), usr.ID, "password")
	require.NoError(t, passwordChangeErr)

	err = db.DeleteUser(context.Background(), usr.ID)
	require.NoError(t, err)
}

func TestLogin(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	var username = "test111"
	var password = "test111"

	_, err = db.CreateUser(context.Background(), username, password)
	require.NoError(t, err)

	_, err = db.Login(context.Background(), "test111", "test111")
	require.NoError(t, err)

	usr, err := db.GetUserByName(context.Background(), "test111")
	require.NoError(t, err)

	err = db.DeleteUser(context.Background(), usr.ID)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	deleteUserErr := db.DeleteUser(context.Background(), "test1")
	require.NoError(t, deleteUserErr)
}
