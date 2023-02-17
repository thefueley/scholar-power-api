package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thefueley/scholar-power-api/internal/exercise"
)

type ExerciseService interface {
	GetExerciseByID(ctx context.Context, uid string) (exercise.Exercise, error)
	GetExerciseByName(ctx context.Context, name string) ([]exercise.Exercise, error)
	GetExerciseByMuscle(ctx context.Context, name string) ([]exercise.Exercise, error)
	GetExerciseByEquipment(ctx context.Context, equip string) ([]exercise.Exercise, error)
}

type ExerciseResponse struct {
	Message string
}

type ExerciseRequest struct {
	Name      string `json:"name"`
	Muscle    string `json:"muscle"`
	Equipment string `json:"equipment"`
}

func (h *SwoleHandler) GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	exercise, err := h.EService.GetExerciseByID(r.Context(), uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(exercise); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetExerciseByName(w http.ResponseWriter, r *http.Request) {
	var req ExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding GetExerciseByName request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedExercise := exerciseRequestToExercise(req)

	exercise, err := h.EService.GetExerciseByName(r.Context(), convertedExercise.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(exercise); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetExerciseByMuscle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	muscleGroup := vars["muscle"]

	if muscleGroup == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	exercises, err := h.EService.GetExerciseByMuscle(r.Context(), muscleGroup)
	if err != nil {
		http.Error(w, "exercises for muscle group not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(exercises); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetExerciseByEquipment(w http.ResponseWriter, r *http.Request) {
	var req ExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding GetExerciseByEquipment request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedExercise := exerciseRequestToExercise(req)

	exercise, err := h.EService.GetExerciseByEquipment(r.Context(), convertedExercise.Equipment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(exercise); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func exerciseRequestToExercise(req ExerciseRequest) exercise.Exercise {
	return exercise.Exercise{
		Name:      req.Name,
		Muscle:    req.Muscle,
		Equipment: req.Equipment,
	}
}
