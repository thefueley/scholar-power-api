package exercise

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrExerciseNotFound = errors.New("exercise not found")
)

type Exercise struct {
	ID           string
	Name         string
	Muscle       string
	Equipment    string
	Instructions string
}

type ExerciseStore interface {
	GetByID(context.Context, string) (Exercise, error)
	GetByName(context.Context, string) (Exercise, error)
	GetByMuscle(context.Context, string) (Exercise, error)
	GetByEquipment(context.Context, string) (Exercise, error)
}

type ExerciseService struct {
	Store ExerciseStore
}

func NewExerciseService(store ExerciseStore) *ExerciseService {
	return &ExerciseService{
		Store: store,
	}
}

func (es *ExerciseService) GetByID(ctx context.Context, id string) (Exercise, error) {
	ex, err := es.Store.GetByID(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by id", err)
		return Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}

func (es *ExerciseService) GetByName(ctx context.Context, id string) (Exercise, error) {
	ex, err := es.Store.GetByName(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by name", err)
		return Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}

func (es *ExerciseService) GetByMuscle(ctx context.Context, id string) (Exercise, error) {
	ex, err := es.Store.GetByMuscle(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by muscle", err)
		return Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}

func (es *ExerciseService) GetByEquipment(ctx context.Context, id string) (Exercise, error) {
	ex, err := es.Store.GetByEquipment(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by equipment", err)
		return Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}
