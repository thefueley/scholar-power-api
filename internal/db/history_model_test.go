// build +integration
package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thefueley/scholar-power-api/internal/history"
)

func TestCreateHistory(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	hist := history.History{
		ID:        "1",
		Date:      "2021-01-01",
		Duration:  "1 hour",
		Notes:     "notes",
		PlanID:    "1",
		AthleteID: "1",
	}

	err = db.CreateHistory(context.Background(), hist)

	require.NoError(t, err)
}

func TestGetHistory(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	exr, err := db.GetExerciseByID(context.Background(), "1")
	require.NoError(t, err)

	require.NotEmpty(t, exr)

	require.NotZero(t, exr.ID)
	require.Equal(t, exr.ID, "1")
}

func TestUpdateHistory(t *testing.T) {
	db, err := NewDatabase("")
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

func TestDeleteHistory(t *testing.T) {
	db, err := NewDatabase("")
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
