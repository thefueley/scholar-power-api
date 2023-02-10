// build +integration
package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thefueley/scholar-power-api/internal/workout"
)

func TestCreateWorkout(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	testWorkout := workout.Workout{
		PlanID:     "1",
		Name:       "Test Workout",
		Sets:       "1",
		Reps:       "1",
		CreatorID:  "1",
		ExerciseID: "1",
	}

	createWorkoutErr := db.CreateWorkout(context.Background(), testWorkout)
	require.NoError(t, createWorkoutErr)
}

func TestGetWorkoutByID(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	wo, err := db.GetWorkoutByID(context.Background(), "1")
	require.NoError(t, err)

	require.NotEmpty(t, wo)

	for _, v := range wo {
		if v.Name == "Test Workout" {
			require.Equal(t, v.Name, "Test Workout")
		}
	}
}

func TestGetWorkoutByUser(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	wo, err := db.GetWorkoutByUser(context.Background(), "test")
	require.NoError(t, err)
	require.NotEmpty(t, wo)
	for _, v := range wo {
		if v.Name == "Test Workout" {
			require.Equal(t, v.Name, "Test Workout")
		}
	}
}
