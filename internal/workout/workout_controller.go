package workout

import (
	"context"
	"errors"
	"fmt"
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
	PlanID    string
	Name      string
	CreatedAt string
	EditedAt  string
	CreatorID string
}

type WorkoutStore interface {
	CreateWorkout(context.Context, Workout) error
	GetWorkoutByID(context.Context, string) ([]WorkoutRow, error)
	GetWorkoutDetails(context.Context, string) ([]Workout, error)
	GetWorkoutByUser(context.Context, string) ([]WorkoutShortInfo, error)
	UpdateWorkout(context.Context, Workout) error
	DeleteWorkout(context.Context, string) error
}

type WorkoutService struct {
	Store WorkoutStore
}

func NewWorkoutService(store WorkoutStore) *WorkoutService {
	return &WorkoutService{
		Store: store,
	}
}

func (ws *WorkoutService) CreateWorkout(ctx context.Context, wo Workout) error {
	err := ws.Store.CreateWorkout(ctx, wo)
	if err != nil {
		fmt.Println("controller.CreateWorkout: ", err)
		return err
	}
	return nil
}

func (ws *WorkoutService) GetWorkoutByID(ctx context.Context, plan_id string) ([]WorkoutRow, error) {
	wo, err := ws.Store.GetWorkoutByID(ctx, plan_id)
	if err != nil {
		fmt.Println("controller.GetWorkoutByID: ", err)
		return []WorkoutRow{}, ErrWorkoutNotFound
	}
	return wo, nil
}

func (ws *WorkoutService) GetWorkoutDetails(ctx context.Context, id string) ([]Workout, error) {
	wo, err := ws.Store.GetWorkoutDetails(ctx, id)
	if err != nil {
		fmt.Println("controller.GetWorkoutDetails: ", err)
		return []Workout{}, ErrWorkoutNotFound
	}
	return wo, nil
}

func (ws *WorkoutService) GetWorkoutByUser(ctx context.Context, user string) ([]WorkoutShortInfo, error) {
	wo, err := ws.Store.GetWorkoutByUser(ctx, user)
	if err != nil {
		fmt.Println("controller.GetWorkoutByUser: ", err)
		return []WorkoutShortInfo{}, ErrWorkoutNotFound
	}
	return wo, nil
}

func (ws *WorkoutService) UpdateWorkout(ctx context.Context, wo Workout) error {
	err := ws.Store.UpdateWorkout(ctx, wo)
	if err != nil {
		fmt.Println("controller.UpdateWorkout: ", err)
		return err
	}
	return nil
}

func (ws *WorkoutService) DeleteWorkout(ctx context.Context, id string) error {
	err := ws.Store.DeleteWorkout(ctx, id)
	if err != nil {
		fmt.Println("controller.DeleteWorkout: ", err)
		return err
	}
	return nil
}
