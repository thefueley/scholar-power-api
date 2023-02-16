package workout

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrWorkoutNotFound = errors.New("workout not found")
)

type Workout struct {
	ID             string
	PlanID         string
	Name           string
	Sets           string
	Reps           string
	Load           string
	CreatedAt      string
	EditedAt       string
	CreatorID      string
	ExerciseID     string
	InstructionsID string
}

type WorkoutRow struct {
	ID                   string
	PlanID               string
	Name                 string
	Sets                 string
	Reps                 string
	Load                 string
	ExerciseName         string
	ExerciseMuscle       string
	ExerciseEquipment    string
	ExerciseInstructions string
}

type WorkoutShortInfo struct {
	PlanID    uuid.UUID
	Name      string
	CreatedAt string
	EditedAt  string
	CreatorID string
}

type WorkoutStore interface {
	CreateWorkout(context.Context, []Workout) error
	GetWorkoutExercises(context.Context, string) ([]WorkoutRow, error)
	GetWorkoutDetails(context.Context, string) ([]Workout, error)
	GetWorkoutsByUser(context.Context, string) ([]WorkoutShortInfo, error)
	UpdateWorkout(context.Context, []Workout) error
	DeleteWorkout(context.Context, []string) error
}

type WorkoutService struct {
	Store WorkoutStore
}

func NewWorkoutService(store WorkoutStore) *WorkoutService {
	return &WorkoutService{
		Store: store,
	}
}

func (ws *WorkoutService) CreateWorkout(ctx context.Context, wo []Workout) error {
	err := ws.Store.CreateWorkout(ctx, wo)
	if err != nil {
		fmt.Println("controller.CreateWorkout: ", err)
		return err
	}
	return nil
}

func (ws *WorkoutService) GetWorkoutExercises(ctx context.Context, plan_id string) ([]WorkoutRow, error) {
	wo, err := ws.Store.GetWorkoutExercises(ctx, plan_id)
	if err != nil {
		fmt.Println("controller.GetWorkoutByID: ", err)
		return []WorkoutRow{}, ErrWorkoutNotFound
	}
	return wo, nil
}

func (ws *WorkoutService) GetWorkoutDetails(ctx context.Context, plan_id string) ([]Workout, error) {
	wo, err := ws.Store.GetWorkoutDetails(ctx, plan_id)
	if err != nil {
		fmt.Println("controller.GetWorkoutDetails: ", err)
		return []Workout{}, ErrWorkoutNotFound
	}
	return wo, nil
}

func (ws *WorkoutService) GetWorkoutsByUser(ctx context.Context, user string) ([]WorkoutShortInfo, error) {
	wo, err := ws.Store.GetWorkoutsByUser(ctx, user)
	if err != nil {
		fmt.Println("controller.GetWorkoutByUser: ", err)
		return []WorkoutShortInfo{}, ErrWorkoutNotFound
	}
	return wo, nil
}

func (ws *WorkoutService) UpdateWorkout(ctx context.Context, wo []Workout) error {
	err := ws.Store.UpdateWorkout(ctx, wo)
	if err != nil {
		fmt.Println("controller.UpdateWorkout: ", err)
		return err
	}
	return nil
}

func (ws *WorkoutService) DeleteWorkout(ctx context.Context, id []string) error {
	err := ws.Store.DeleteWorkout(ctx, id)
	if err != nil {
		fmt.Println("controller.DeleteWorkout: ", err)
		return err
	}
	return nil
}
