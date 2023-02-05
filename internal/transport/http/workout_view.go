package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thefueley/scholar-power-api/internal/workout"
)

type WorkoutService interface {
	CreateWorkout(ctx context.Context, wo workout.Workout) error
	GetWorkoutByID(ctx context.Context, id string) ([]workout.Workout, error)
	GetWorkoutByUser(ctx context.Context, user string) ([]workout.Workout, error)
	UpdateWorkout(ctx context.Context, wo workout.Workout) error
	DeleteWorkout(ctx context.Context, id string) error
}

type WorkoutResponse struct {
	Message string
}

type WorkoutRequest struct {
	ID         string `json:"id"`
	WorkoutID  string `json:"workout_id"`
	Name       string `json:"name"`
	Sets       string `json:"sets"`
	Reps       string `json:"reps"`
	CreatedAt  string `json:"created_at"`
	CreatorID  string `json:"creator_id"`
	ExerciseID string `json:"exercise_id"`
}

func (h *SwoleHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var req WorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding CreateWorkout request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedWorkout := workoutRequestToWorkout(req)

	if err := h.WService.CreateWorkout(r.Context(), convertedWorkout); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(WorkoutResponse{Message: "workout created"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wid := vars["workout_id"]

	workout, err := h.WService.GetWorkoutByID(r.Context(), wid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetWorkoutByUser(w http.ResponseWriter, r *http.Request) {
	var req WorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding GetWorkoutByUser request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	workout, err := h.WService.GetWorkoutByUser(r.Context(), req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	var req WorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding UpdateWorkout request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedWorkout := workoutRequestToWorkout(req)

	if err := h.WService.UpdateWorkout(r.Context(), convertedWorkout); err != nil {
		fmt.Printf("view.UpdateWorkout: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(WorkoutResponse{Message: "workout updated"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wid := vars["workout_id"]

	if err := h.WService.DeleteWorkout(r.Context(), wid); err != nil {
		fmt.Printf("view.DeleteWorkout: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(WorkoutResponse{Message: "workout deleted"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func workoutRequestToWorkout(req WorkoutRequest) workout.Workout {
	return workout.Workout{
		ID:         req.ID,
		WorkoutID:  req.WorkoutID,
		Name:       req.Name,
		Sets:       req.Sets,
		Reps:       req.Reps,
		CreatedAt:  req.CreatedAt,
		CreatorID:  req.CreatorID,
		ExerciseID: req.ExerciseID,
	}
}
