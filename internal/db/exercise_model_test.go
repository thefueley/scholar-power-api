package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetExerciseByID(t *testing.T) {
	db, err := NewDatabase()
	require.NoError(t, err)

	exr, err := db.GetExerciseByID(context.Background(), "1")
	require.NoError(t, err)

	require.NotEmpty(t, exr)

	require.NotZero(t, exr.ID)
	require.Equal(t, exr.ID, "1")
}

func TestGetExerciseByName(t *testing.T) {
	db, err := NewDatabase()
	require.NoError(t, err)

	exr, err := db.GetExerciseByName(context.Background(), "Pushups")
	require.NoError(t, err)
	require.NotEmpty(t, exr)
	for _, v := range exr {
		if v.Name == "Pushups" {
			require.Equal(t, v.Name, "Pushups")
		}
	}
}

func TestGetExerciseByMuscle(t *testing.T) {
	db, err := NewDatabase()
	require.NoError(t, err)

	exr, err := db.GetExerciseByMuscle(context.Background(), "chest")
	require.NoError(t, err)
	require.NotEmpty(t, exr)
	for _, v := range exr {
		if v.Name == "chest" {
			require.Equal(t, v.Name, "chest")
		}
	}
}

func TestGetExerciseByEquipment(t *testing.T) {
	db, err := NewDatabase()
	require.NoError(t, err)

	exr, err := db.GetExerciseByEquipment(context.Background(), "barbell")
	require.NoError(t, err)
	require.NotEmpty(t, exr)
	for _, v := range exr {
		if v.Name == "barbell" {
			require.Equal(t, v.Equipment, "barbell")
		}
	}
}
