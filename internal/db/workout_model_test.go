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
		PlanID:         "999999",
		Name:           "999999",
		Sets:           "999999",
		Reps:           "999999",
		Load:           "999999",
		CreatorID:      "3",
		ExerciseID:     "1",
		InstructionsID: "1",
	}

	createWorkoutErr := db.CreateWorkout(context.Background(), testWorkout)
	require.NoError(t, createWorkoutErr)
}

func TestGetWorkoutExercises(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	wo, err := db.GetWorkoutExercises(context.Background(), "999999")
	require.NoError(t, err)

	require.NotEmpty(t, wo)

	for _, v := range wo {
		if v.Name == "999999" {
			require.Equal(t, v.Name, "999999")
		}
	}
}

func TestGetWorkoutsByUser(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	wo, err := db.GetWorkoutsByUser(context.Background(), "test1")
	require.NoError(t, err)
	require.NotEmpty(t, wo)
	for _, v := range wo {
		if v.Name == "999999" {
			require.Equal(t, v.Name, "999999")
		}
	}
}
