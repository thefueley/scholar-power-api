package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/thefueley/scholar-power-api/internal/exercise"
)

type ExerciseRow struct {
	ID           sql.NullString
	Name         sql.NullString
	Muscle       sql.NullString
	Equipment    sql.NullString
	Instructions sql.NullString
}

func (db *Database) GetExerciseByID(ctx context.Context, id string) (exercise.Exercise, error) {
	var exr ExerciseRow
	row := db.QueryRowContext(ctx,
		`SELECT * 
		FROM exercise
		WHERE id = $1`,
		id,
	)
	err := row.Scan(&exr.ID, &exr.Name, &exr.Muscle, &exr.Equipment, &exr.Instructions)

	if err != nil {
		return exercise.Exercise{}, fmt.Errorf("could not get exercise: %w", err)
	}
	return exerciseRowToExercise(exr), nil
}

func (db *Database) GetExerciseByName(ctx context.Context, name string) ([]exercise.Exercise, error) {
	row, err := db.QueryContext(ctx,
		`SELECT *
		FROM exercise
		WHERE name = $1`,
		name,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundExercises := make([]exercise.Exercise, 0)

	for row.Next() {
		var exercise exercise.Exercise
		if err := row.Scan(&exercise.ID, &exercise.Name, &exercise.Muscle, &exercise.Equipment, &exercise.Instructions); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		foundExercises = append(foundExercises, exercise)
	}

	if err != nil {
		return []exercise.Exercise{}, fmt.Errorf("could not get exercise: %w", err)
	}
	return foundExercises, nil
}

func (db *Database) GetExerciseByMuscle(ctx context.Context, muscle string) ([]exercise.Exercise, error) {
	row, err := db.QueryContext(ctx,
		`SELECT *
		FROM exercise
		WHERE muscle = $1`,
		muscle,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundExercises := make([]exercise.Exercise, 0)

	for row.Next() {
		var exercise exercise.Exercise
		if err := row.Scan(&exercise.ID, &exercise.Name, &exercise.Muscle, &exercise.Equipment, &exercise.Instructions); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		foundExercises = append(foundExercises, exercise)
	}

	if err != nil {
		return []exercise.Exercise{}, fmt.Errorf("could not get exercise: %w", err)
	}
	return foundExercises, nil
}

func (db *Database) GetExerciseByEquipment(ctx context.Context, equipment string) ([]exercise.Exercise, error) {
	row, err := db.QueryContext(ctx,
		`SELECT *
		FROM exercise
		WHERE equipment = $1`,
		equipment,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	foundExercises := make([]exercise.Exercise, 0)

	for row.Next() {
		var exercise exercise.Exercise
		if err := row.Scan(&exercise.ID, &exercise.Name, &exercise.Muscle, &exercise.Equipment, &exercise.Instructions); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		foundExercises = append(foundExercises, exercise)
	}

	if err != nil {
		return []exercise.Exercise{}, fmt.Errorf("could not get exercise: %w", err)
	}
	return foundExercises, nil
}

func exerciseRowToExercise(exr ExerciseRow) exercise.Exercise {
	return exercise.Exercise{
		ID:           exr.ID.String,
		Name:         exr.Name.String,
		Muscle:       exr.Muscle.String,
		Equipment:    exr.Equipment.String,
		Instructions: exr.Instructions.String,
	}
}
