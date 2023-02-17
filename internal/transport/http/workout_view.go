package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/thefueley/scholar-power-api/internal/workout"
)

type WorkoutService interface {
	CreateWorkout(ctx context.Context, wo []workout.Workout) error
	GetWorkoutExercises(ctx context.Context, id string) ([]workout.WorkoutRow, error)
	GetWorkoutDetails(ctx context.Context, plan_id string) ([]workout.Workout, error)
	GetWorkoutsByUser(ctx context.Context, user string) ([]workout.WorkoutShortInfo, error)
	UpdateWorkout(ctx context.Context, wo []workout.Workout) error
	DeleteWorkout(ctx context.Context, id []string) error
}

type WorkoutResponse struct {
	Message string
}

type WorkoutRequest struct {
	ID             string `json:"id"`
	PlanID         string `json:"plan_id"`
	Name           string `json:"name"`
	Sets           string `json:"sets"`
	Reps           string `json:"reps"`
	Load           string `json:"load"`
	CreatedAt      string `json:"created_at"`
	EditedAt       string `json:"edited_at"`
	CreatorID      string `json:"creator_id"`
	ExerciseID     string `json:"exercise_id"`
	InstructionsID string `json:"instructions_id"`
}

type WorkoutLineItem struct {
	Sets       string `json:"sets"`
	Reps       string `json:"reps"`
	Load       string `json:"load"`
	ExerciseID string `json:"exercise_id"`
}

type WorkoutPlanRequest struct {
	UID       string            `json:"uid"`
	Name      string            `json:"name"`
	Exercises []WorkoutLineItem `json:"exercises"`
}

func (h *SwoleHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var req WorkoutPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding CreateWorkout request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userWorkoutPlan := workoutPlanRequestToWorkout(req)

	err := h.AuthZ(r, userWorkoutPlan[0].CreatorID)
	if err != nil {
		fmt.Printf("view.CreateWorkout AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.WService.CreateWorkout(r.Context(), userWorkoutPlan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(WorkoutResponse{Message: "workout created"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetWorkoutExercises(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid := vars["plan_id"]

	wod, err := h.WService.GetWorkoutDetails(r.Context(), pid)
	if err != nil {
		fmt.Println("view.GetWorkoutExercises: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.AuthZ(r, wod[0].CreatorID)
	if err != nil {
		fmt.Printf("view.GetWorkoutExercises AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	workout, err := h.WService.GetWorkoutExercises(r.Context(), pid)
	if err != nil {
		fmt.Println("view.GetWorkoutExercises: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetWorkoutsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	uid, err := h.UService.GetUserByName(r.Context(), username)
	if err != nil {
		fmt.Println("view.GetWorkoutsByUser: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if uid.ID == "" {
		fmt.Println("view.GetWorkoutsByUser: creator_id is empty")
		http.Error(w, "creator_id is empty", http.StatusBadRequest)
		return
	}

	err = h.AuthZ(r, uid.ID)
	if err != nil {
		fmt.Printf("view.GetWorkoutsByUser AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	workout, err := h.WService.GetWorkoutsByUser(r.Context(), username)
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
	vars := mux.Vars(r)
	pid := vars["plan_id"]

	var updateWorkoutRequest []WorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&updateWorkoutRequest); err != nil {
		fmt.Printf("error decoding UpdateWorkout request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if pid == "" {
		fmt.Println("view.UpdateWorkout: plan_id is empty")
		http.Error(w, "plan_id is empty", http.StatusBadRequest)
		return
	}

	allUpdateWorkoutRequestItems := make([]workout.Workout, 0)

	for wo := range updateWorkoutRequest {
		oneWorkout := workoutRequestToWorkout(updateWorkoutRequest[wo])
		oneWorkout.PlanID = pid
		allUpdateWorkoutRequestItems = append(allUpdateWorkoutRequestItems, oneWorkout)
	}

	rid, err := h.WService.GetWorkoutDetails(r.Context(), pid)
	if err != nil {
		fmt.Println("view.UpdateWorkout: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.AuthZ(r, rid[0].CreatorID)
	if err != nil {
		fmt.Printf("view.UpdateWorkout AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.WService.UpdateWorkout(r.Context(), allUpdateWorkoutRequestItems); err != nil {
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
	pid := vars["plan_id"]

	if pid == "" {
		fmt.Println("view.UpdateWorkout: plan_id is empty")
		http.Error(w, "plan_id is empty", http.StatusBadRequest)
		return
	}

	workoutsInPlan, err := h.WService.GetWorkoutDetails(r.Context(), pid)
	if err != nil {
		fmt.Printf("view.DeleteWorkout GetWorkoutDetails: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.AuthZ(r, workoutsInPlan[0].CreatorID)
	if err != nil {
		fmt.Printf("view.DeleteWorkout AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	wid := make([]string, 0)
	for _, v := range workoutsInPlan {
		wid = append(wid, v.ID)
	}

	if err := h.WService.DeleteWorkout(r.Context(), wid); err != nil {
		fmt.Printf("view.DeleteWorkout: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(WorkoutResponse{Message: "Poof! It's gone."}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func workoutRequestToWorkout(req WorkoutRequest) workout.Workout {
	return workout.Workout{
		ID:             req.ID,
		PlanID:         req.PlanID,
		Name:           req.Name,
		Sets:           req.Sets,
		Reps:           req.Reps,
		Load:           req.Load,
		CreatedAt:      req.CreatedAt,
		EditedAt:       req.EditedAt,
		CreatorID:      req.CreatorID,
		ExerciseID:     req.ExerciseID,
		InstructionsID: req.InstructionsID,
	}
}

func workoutPlanRequestToWorkout(req WorkoutPlanRequest) []workout.Workout {

	var UID = req.UID
	var PlanID = uuid.New().String()
	var Name = req.Name
	now := time.Now().Format("01-02-2006 15:04")

	workoutInfo := make([]workout.Workout, 0)

	for _, v := range req.Exercises {
		workoutInfo = append(workoutInfo, workout.Workout{
			PlanID:         PlanID,
			Name:           Name,
			Sets:           v.Sets,
			Reps:           v.Reps,
			Load:           v.Load,
			CreatedAt:      now,
			EditedAt:       now,
			CreatorID:      UID,
			ExerciseID:     v.ExerciseID,
			InstructionsID: v.ExerciseID,
		})
	}
	return workoutInfo
}
