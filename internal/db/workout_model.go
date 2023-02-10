package db

import (
	"context"
	"fmt"
	"log"

	"github.com/thefueley/scholar-power-api/internal/workout"
)

func (db *Database) CreateWorkout(ctx context.Context, workout workout.Workout) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO workout_plan (workout_id, name, sets, reps, load, created_at, creator_id, exercise_id, instructions_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		workout.PlanID,
		workout.Name,
		workout.Sets,
		workout.Reps,
		workout.Load,
		workout.CreatedAt,
		workout.CreatorID,
		workout.ExerciseID,
		workout.InstructionsID,
	)

	if err != nil {
		return fmt.Errorf("could not create workout: %w", err)
	}

	return nil
}

func (db *Database) GetWorkoutByID(ctx context.Context, plan_id string) ([]workout.WorkoutRow, error) {
	row, err := db.QueryContext(ctx,
		`SELECT 
		workout.id, 
		workout.plan_id, 
		workout.name, 
		workout.sets, 
		workout.reps, 
		workout.load, 
		exercise.name, 
		exercise.muscle, 
		exercise.equipment, 
		exercise.instructions 
		FROM workout 
		JOIN exercise ON workout.exercise_id = exercise.id 
		WHERE workout.plan_id = $1;`,
		plan_id,
	)

	if err != nil {
		fmt.Println("model.GetWorkoutByID: QueryContext: ", err.Error())
		log.Fatal(err)
	}

	defer row.Close()

	foundWorkouts := make([]workout.WorkoutRow, 0)

	for row.Next() {
		var wo workout.WorkoutRow
		if err := row.Scan(&wo.ID, &wo.PlanID, &wo.Name, &wo.Sets, &wo.Reps, &wo.Load, &wo.ExerciseName, &wo.ExerciseMuscle, &wo.ExerciseEquipment, &wo.ExerciseInstructions); err != nil {
			fmt.Println("model.GetWorkoutByID: Scan: ", err.Error())
			log.Fatal(err)
		}
		foundWorkouts = append(foundWorkouts, wo)
	}

	if err != nil {
		return []workout.WorkoutRow{}, fmt.Errorf("could not get workout: %w", err)
	}
	return foundWorkouts, nil
}

func (db *Database) GetWorkoutDetails(ctx context.Context, id string) ([]workout.Workout, error) {
	row, err := db.QueryContext(ctx,
		`SELECT 
		id, 
		plan_id, 
		name, 
		sets, 
		reps, 
		load, 
		created_at, 
		edited_at, 
		creator_id, 
		exercise_id, 
		instructions_id
		FROM workout 
		WHERE id = $1`,
		id,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundWorkouts := make([]workout.Workout, 0)

	for row.Next() {
		var wo workout.Workout
		if err := row.Scan(&wo.ID, &wo.PlanID, &wo.Name, &wo.Sets, &wo.Reps, &wo.Load, &wo.CreatedAt, &wo.EditedAt, &wo.CreatorID, &wo.ExerciseID, &wo.InstructionsID); err != nil {
			log.Fatal(err)
		}
		foundWorkouts = append(foundWorkouts, wo)
	}

	if err != nil {
		return []workout.Workout{}, fmt.Errorf("could not get workout: %w", err)
	}
	return foundWorkouts, nil
}

func (db *Database) GetWorkoutByUser(ctx context.Context, user string) ([]workout.WorkoutShortInfo, error) {
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
		`SELECT plan_id, name, created_at, edited_at, creator_id 
		FROM workout WHERE creator_id = $1 
		GROUP BY name`,
		userID,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	foundWorkouts := make([]workout.WorkoutShortInfo, 0)

	for row.Next() {
		var wo workout.WorkoutShortInfo
		if err := row.Scan(&wo.PlanID, &wo.Name, &wo.CreatedAt, &wo.EditedAt, &wo.CreatorID); err != nil {
			log.Fatal(err)
		}

		foundWorkouts = append(foundWorkouts, wo)
	}

	if err != nil {
		return []workout.WorkoutShortInfo{}, fmt.Errorf("could not find workout: %w", err)
	}
	return foundWorkouts, nil
}

func (db *Database) UpdateWorkout(ctx context.Context, workout workout.Workout) error {
	_, err := db.ExecContext(ctx,
		`UPDATE workout_plan
		SET name = $1, sets = $2, reps = $3, exercise_id = $4
		WHERE id = $5`,
		workout.Name,
		workout.Sets,
		workout.Reps,
		workout.ExerciseID,
		workout.ID,
	)

	if err != nil {
		return fmt.Errorf("could not update workout: %w", err)
	}

	return nil
}

func (db *Database) DeleteWorkout(ctx context.Context, id string) error {
	_, err := db.ExecContext(ctx,
		`DELETE FROM workout_plan
		WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("could not delete workout: %w", err)
	}

	return nil
}
