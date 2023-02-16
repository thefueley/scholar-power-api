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
	GetExerciseByID(context.Context, string) (Exercise, error)
	GetExerciseByName(context.Context, string) ([]Exercise, error)
	GetExerciseByMuscle(context.Context, string) ([]Exercise, error)
	GetExerciseByEquipment(context.Context, string) ([]Exercise, error)
}

type ExerciseService struct {
	Store ExerciseStore
}

func NewExerciseService(store ExerciseStore) *ExerciseService {
	return &ExerciseService{
		Store: store,
	}
}

func (es *ExerciseService) GetExerciseByID(ctx context.Context, id string) (Exercise, error) {
	ex, err := es.Store.GetExerciseByID(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by id", err)
		return Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}

func (es *ExerciseService) GetExerciseByName(ctx context.Context, id string) ([]Exercise, error) {
	ex, err := es.Store.GetExerciseByName(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by name:", err)
		return []Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}

func (es *ExerciseService) GetExerciseByMuscle(ctx context.Context, id string) ([]Exercise, error) {
	ex, err := es.Store.GetExerciseByMuscle(ctx, id)
	if err != nil {
		fmt.Println("error getting exercise by muscle: ", err)
		return []Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}

func (es *ExerciseService) GetExerciseByEquipment(ctx context.Context, equip string) ([]Exercise, error) {
	ex, err := es.Store.GetExerciseByEquipment(ctx, equip)
	if err != nil {
		fmt.Println("error getting exercise by equipment", err)
		return []Exercise{}, ErrExerciseNotFound
	}
	return ex, nil
}
