package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thefueley/scholar-power-api/internal/history"
)

type HistoryService interface {
	CreateHistory(context.Context, history.History) error
	GetHistory(context.Context, string) ([]history.History, error)
	UpdateHistory(context.Context, history.History) error
	DeleteHistory(context.Context, string) error
}

type HistoryResponse struct {
	Message string
}

type HistoryRequest struct {
	ID        string
	Date      string `json:"date"`
	Duration  string `json:"duration"`
	Notes     string `json:"notes"`
	PlanID    string `json:"plan_id"`
	AthleteID string `json:"athlete_id"`
}

func (h *SwoleHandler) CreateHistory(w http.ResponseWriter, r *http.Request) {
	var req HistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding view.CreateHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.AuthZ(r, req.AthleteID)
	if err != nil {
		fmt.Printf("view.CreateHistory AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	convertedHistory := historyRequestToHistory(req)

	err = h.HService.CreateHistory(r.Context(), convertedHistory)
	if err != nil {
		fmt.Printf("error encoding view.CreateHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "History created"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *SwoleHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	err := h.AuthZ(r, uid)
	if err != nil {
		fmt.Printf("view.GetHistory AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	hist, err := h.HService.GetHistory(r.Context(), uid)
	if err != nil {
		fmt.Printf("error getting view.GetHistoryByUserID request: %v\n", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(hist); err != nil {
		fmt.Printf("error encoding view.GetHistoryByUserID request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req HistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding view.UpdateHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedHistory := historyRequestToHistory(req)
	convertedHistory.ID = id
	convertedHistory.AthleteID = req.AthleteID

	err := h.AuthZ(r, req.AthleteID)
	if err != nil {
		fmt.Printf("view.UpdateHistory AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.HService.UpdateHistory(r.Context(), convertedHistory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(HistoryResponse{Message: "workout history updated"}); err != nil {
		fmt.Printf("error encoding view.UpdateHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid := vars["id"]

	var req HistoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("error decoding view.DeleteHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.AuthZ(r, req.AthleteID)
	if err != nil {
		fmt.Printf("view.GetHistory AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.HService.DeleteHistory(r.Context(), pid)
	fmt.Printf("view.DeleteHistory err: %v\n", err)
	if err != nil {
		fmt.Printf("error deleting view.DeleteHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(HistoryResponse{Message: "Poof! It's gone."}); err != nil {
		fmt.Printf("error encoding view.DeleteHistory request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func historyRequestToHistory(req HistoryRequest) history.History {
	return history.History{
		ID:        req.ID,
		Date:      req.Date,
		Duration:  req.Duration,
		Notes:     req.Notes,
		PlanID:    req.PlanID,
		AthleteID: req.AthleteID,
	}
}
