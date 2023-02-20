package db

import (
	"context"
	"fmt"
	"log"

	"github.com/thefueley/scholar-power-api/internal/workout"
)

func (db *Database) CreateWorkout(ctx context.Context, workout []workout.Workout) error {
	for _, wo := range workout {
		_, err := db.ExecContext(ctx,
			`INSERT INTO workout (plan_id, name, sets, reps, load, created_at, edited_at, creator_id, exercise_id, instructions_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			wo.PlanID,
			wo.Name,
			wo.Sets,
			wo.Reps,
			wo.Load,
			wo.CreatedAt,
			wo.EditedAt,
			wo.CreatorID,
			wo.ExerciseID,
			wo.InstructionsID,
		)

		if err != nil {
			return fmt.Errorf("could not create workout: %w", err)
		}
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
	}

	defer row.Close()

	foundWorkouts := make([]workout.WorkoutRow, 0)

	for row.Next() {
		var wo workout.WorkoutRow
		if err := row.Scan(&wo.ID, &wo.PlanID, &wo.Name, &wo.Sets, &wo.Reps, &wo.Load, &wo.ExerciseName, &wo.ExerciseMuscle, &wo.ExerciseEquipment, &wo.ExerciseInstructions); err != nil {
			fmt.Println("model.GetWorkoutByID: Scan: ", err.Error())
			// log.Fatal(err)
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
		*
		FROM workout 
		WHERE plan_id = $1`,
		plan_id,
	)

	if err != nil {
		return []workout.Workout{}, fmt.Errorf("model.GetWorkoutDetails: QueryContext: %v", err)
	}

	defer row.Close()

	foundWorkouts := make([]workout.Workout, 0)

	for row.Next() {
		var wo workout.Workout
		if err := row.Scan(&wo.ID, &wo.PlanID, &wo.Name, &wo.Sets, &wo.Reps, &wo.Load, &wo.CreatedAt, &wo.EditedAt, &wo.CreatorID, &wo.ExerciseID, &wo.InstructionsID); err != nil {
			fmt.Printf("model.GetWorkoutDetails: row.Scan: %v\n", err)
		}
		foundWorkouts = append(foundWorkouts, wo)
	}

	if len(foundWorkouts) == 0 {
		return []workout.Workout{}, fmt.Errorf("model.GetWorkoutDetails: no workouts found")
	}

	if err != nil {
		return []workout.Workout{}, fmt.Errorf("model.GetWorkoutDetails: row.Next: %v", err)
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
	// get current workouts
	currentWorkouts, err := db.GetWorkoutDetails(ctx, updatedWorkouts[0].PlanID)

	if err != nil {
		return fmt.Errorf("could not update workout: %w", err)
	}

	// delete workouts that are not in updatedWorkouts
	deleteIDs := selectDeletableLineItems(currentWorkouts, updatedWorkouts)
	for _, id := range deleteIDs {
		_, err := db.ExecContext(ctx, `DELETE FROM workout WHERE id = $1`, id)
		if err != nil {
			return fmt.Errorf("could not update workout: %w", err)
		}
	}

	for k := range updatedWorkouts {
		if updatedWorkouts[k].ID == "" {
			newWorkout := make([]workout.Workout, 0)
			newWorkout = append(newWorkout, updatedWorkouts[k])
			err := db.CreateWorkout(ctx, newWorkout)

			if err != nil {
				return fmt.Errorf("could not update workout: %w", err)
			}
		} else {
			_, err := db.ExecContext(ctx,
				`UPDATE workout SET
			id = $1,
			plan_id = $2, 
			name = $3, 
			sets = $4, 
			reps = $5, 
			load = $6, 
			edited_at = $7, 
			creator_id = $8,
			exercise_id = $9, 
			instructions_id = $10 
			WHERE id = $1`,
				updatedWorkouts[k].ID,
				updatedWorkouts[k].PlanID,
				updatedWorkouts[k].Name,
				updatedWorkouts[k].Sets,
				updatedWorkouts[k].Reps,
				updatedWorkouts[k].Load,
				updatedWorkouts[k].EditedAt,
				updatedWorkouts[k].CreatorID,
				updatedWorkouts[k].ExerciseID,
				updatedWorkouts[k].InstructionsID,
			)

			if err != nil {
				return fmt.Errorf("could not update workout: %w", err)
			}
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

func selectDeletableLineItems(current, updated []workout.Workout) []string {
	var deletableIDs []string

	for _, v := range current {
		var found bool
		for _, w := range updated {
			if v.ID == w.ID {
				found = true
			}
		}
		if !found {
			deletableIDs = append(deletableIDs, v.ID)
		}
	}

	return deletableIDs
}
