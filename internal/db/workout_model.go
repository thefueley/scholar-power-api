package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/thefueley/scholar-power-api/internal/workout"
)

type WorkoutRow struct {
	ID         string
	WorkoutID  sql.NullString
	Name       sql.NullString
	Sets       sql.NullString
	Reps       sql.NullString
	CreatedAt  sql.NullString
	CreatorID  sql.NullString
	ExerciseID sql.NullString
}

func (db *Database) CreateWorkout(ctx context.Context, workout workout.Workout) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO workout_plan (workout_id, name, sets, reps, created_at, creator_id, exercise_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		workout.WorkoutID,
		workout.Name,
		workout.Sets,
		workout.Reps,
		workout.CreatedAt,
		workout.CreatorID,
		workout.ExerciseID,
	)

	if err != nil {
		return fmt.Errorf("could not create workout: %w", err)
	}

	return nil
}

func (db *Database) GetWorkoutByID(ctx context.Context, id string) ([]workout.Workout, error) {
	row, err := db.QueryContext(ctx,
		`SELECT id, workout_id, name, sets, reps, created_at, creator_id, exercise_id
		FROM workout_plan
		WHERE workout_id = $1`,
		id,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundWorkouts := make([]workout.Workout, 0)

	for row.Next() {
		var workout workout.Workout
		if err := row.Scan(&workout.ID, &workout.WorkoutID, &workout.Name, &workout.Sets, &workout.Reps, &workout.CreatedAt, &workout.CreatorID, &workout.ExerciseID); err != nil {
			log.Fatal(err)
		}
		foundWorkouts = append(foundWorkouts, workout)
	}

	if err != nil {
		return []workout.Workout{}, fmt.Errorf("could not get workout: %w", err)
	}
	return foundWorkouts, nil
}

func (db *Database) GetWorkoutByUser(ctx context.Context, user string) ([]workout.Workout, error) {
	userRow := db.QueryRowContext(ctx,
		`SELECT id
		FROM user
		WHERE username = $1`, user)

	var userID string
	err := userRow.Scan(&userID)
	if err != nil {
		fmt.Printf("could not find user: %v in model.GetWorkoutByUser\n", err)
	}

	row, err := db.QueryContext(ctx,
		`SELECT id, workout_id, name, sets, reps, created_at, creator_id, exercise_id
		FROM workout_plan
		WHERE creator_id = $1`,
		userID,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundWorkouts := make([]workout.Workout, 0)

	for row.Next() {
		var wor workout.Workout
		if err := row.Scan(&wor.ID, &wor.WorkoutID, &wor.Name, &wor.Sets, &wor.Reps, &wor.CreatedAt, &wor.CreatorID, &wor.ExerciseID); err != nil {
			log.Fatal(err)
		}
		foundWorkouts = append(foundWorkouts, wor)
	}

	if err != nil {
		return []workout.Workout{}, fmt.Errorf("could not find workout: %w", err)
	}
	return foundWorkouts, nil
}

func (db *Database) UpdateWorkout(ctx context.Context, workout workout.Workout) error {
	_, err := db.ExecContext(ctx,
		`UPDATE workout_plan
		SET name = $1, sets = $2, reps = $3
		WHERE workout_id = $4`,
		workout.Name,
		workout.Sets,
		workout.Reps,
		workout.WorkoutID,
	)

	if err != nil {
		return fmt.Errorf("could not update workout: %w", err)
	}

	return nil
}

func (db *Database) DeleteWorkout(ctx context.Context, id string) error {
	_, err := db.ExecContext(ctx,
		`DELETE FROM workout_plan
		WHERE workout_id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("could not delete workout: %w", err)
	}

	return nil
}

func workoutRowToWorkout(wor WorkoutRow) workout.Workout {
	return workout.Workout{
		ID:         wor.ID,
		WorkoutID:  wor.WorkoutID.String,
		Name:       wor.Name.String,
		Sets:       wor.Sets.String,
		Reps:       wor.Reps.String,
		CreatedAt:  wor.CreatedAt.String,
		CreatorID:  wor.CreatorID.String,
		ExerciseID: wor.ExerciseID.String,
	}
}
