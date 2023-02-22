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
		Name:           "999999",
		Sets:           "999999",
		Reps:           "999999",
		Load:           "999999",
		CreatorID:      "999999",
		ExerciseID:     "1",
		InstructionsID: "1",
	}
	workouts := []workout.Workout{testWorkout}

	createWorkoutErr := db.CreateWorkout(context.Background(), workouts)
	require.NoError(t, createWorkoutErr)

	err = db.DeleteWorkout(context.Background(), []string{"999999"})
	require.NoError(t, err)
}

func TestGetWorkoutExercises(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)

	// create a workout
	testWorkout := workout.Workout{
		PlanID:         "999999",
		Name:           "999999",
		Sets:           "999999",
		Reps:           "999999",
		Load:           "999999",
		CreatorID:      "999999",
		ExerciseID:     "1",
		InstructionsID: "1",
	}
	workouts := []workout.Workout{testWorkout}

	createWorkoutErr := db.CreateWorkout(context.Background(), workouts)
	require.NoError(t, createWorkoutErr)

	// get the workout
	wo, err := db.GetWorkoutExercises(context.Background(), "999999")
	require.NoError(t, err)
	require.NotEmpty(t, wo)

	for _, v := range wo {
		if v.Name == "999999" {
			require.Equal(t, v.Name, "999999")
		}
	}

	// delete the workout
	err = db.DeleteWorkout(context.Background(), []string{"999999"})
	require.NoError(t, err)
}
