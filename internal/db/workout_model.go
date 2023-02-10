package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thefueley/scholar-power-api/internal/workout"
)

func (db *Database) CreateWorkout(ctx context.Context, workout workout.Workout) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO workout (plan_id, name, sets, reps, load, creator_id, exercise_id, instructions_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		workout.PlanID,
		workout.Name,
		workout.Sets,
		workout.Reps,
		workout.Load,
		workout.CreatorID,
		workout.ExerciseID,
		workout.InstructionsID,
	)

	if err != nil {
		return fmt.Errorf("could not create workout: %w", err)
	}

	return nil
}

func (db *Database) GetWorkoutExercises(ctx context.Context, plan_id string) ([]workout.WorkoutRow, error) {
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

func (db *Database) GetWorkoutDetails(ctx context.Context, plan_id string) ([]workout.Workout, error) {
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
		WHERE plan_id = $1`,
		plan_id,
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

func (db *Database) GetWorkoutsByUser(ctx context.Context, user string) ([]workout.WorkoutShortInfo, error) {
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

func (db *Database) UpdateWorkout(ctx context.Context, updatedWorkouts []workout.Workout) error {
	// get current workouts from plan_id
	currentWorkouts, err := db.GetWorkoutDetails(ctx, updatedWorkouts[0].PlanID)
	if err != nil {
		return fmt.Errorf("could not get current workouts: %w", err)
	}

	finalWorkouts := make([]workout.Workout, 0)

	// compare current workouts to new workouts
	// if current workout is outdated, update it
	for _, updated := range updatedWorkouts {
		for _, current := range currentWorkouts {
			if updated.ID == current.ID {
				oneWorkout := reconcileWorkout(current, updated)
				finalWorkouts = append(finalWorkouts, oneWorkout)
			}
		}
	}

	// update workouts
	for k := range finalWorkouts {
		_, err := db.ExecContext(ctx,
			`UPDATE workout SET 
			name = $1, 
			sets = $2, 
			reps = $3, 
			load = $4, 
			edited_at = $5,
			exercise_id = $6, 
			instructions_id = $7 
			WHERE id = $8`,
			finalWorkouts[k].Name,
			finalWorkouts[k].Sets,
			finalWorkouts[k].Reps,
			finalWorkouts[k].Load,
			time.Now(),
			finalWorkouts[k].ExerciseID,
			finalWorkouts[k].InstructionsID,
			finalWorkouts[k].ID,
		)

		if err != nil {
			return fmt.Errorf("could not update workout: %w", err)
		}
	}

	return nil
}

func (db *Database) DeleteWorkout(ctx context.Context, id []string) error {
	for _, v := range id {
		_, err := db.ExecContext(ctx,
			`DELETE FROM workout
			WHERE id = $1`,
			v,
		)

		if err != nil {
			return fmt.Errorf("could not delete workout: %w", err)
		}
	}
	return nil
}

func reconcileWorkout(current, updated workout.Workout) workout.Workout {
	if updated.Name == "" {
		updated.Name = current.Name
	}
	if updated.Sets == "" {
		updated.Sets = current.Sets
	}
	if updated.Reps == "" {
		updated.Reps = current.Reps
	}
	if updated.Load == "" {
		updated.Load = current.Load
	}
	if updated.ExerciseID == "" {
		updated.ExerciseID = current.ExerciseID
	}
	if updated.InstructionsID == "" {
		updated.InstructionsID = current.InstructionsID
	}
	return updated
}
