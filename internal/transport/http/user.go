package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	swoleuser "github.com/thefueley/scholar-power-api/internal/user"
)

type UserService interface {
	CreateUser(ctx context.Context, username string, password string) (swoleuser.User, error)
	GetByID(ctx context.Context, uid string) (swoleuser.User, error)
	GetByUserName(ctx context.Context, username string) (swoleuser.User, error)
	UpdateUserPassword(ctx context.Context, uid string, password string) error
	DeleteUser(ctx context.Context, uid string) error
}

type Response struct {
	Message string
}

type CreateUserRequest struct {
	UserName     string `json:"username" validate:"required"`
	PasswordHash string `json:"password_hash" validate:"required"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		http.Error(w, "not a valid user", http.StatusBadRequest)
		return
	}

	convertedUser := createUserRequestToUser(user)

	createdUser, err := h.Service.CreateUser(r.Context(), convertedUser.UserName, convertedUser.PasswordHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := h.Service.GetByID(r.Context(), uid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(usr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetByUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uname := vars["username"]
	if uname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := h.Service.GetByUserName(r.Context(), uname)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(usr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user swoleuser.User
	user.ID = uid
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.UpdateUserPassword(r.Context(), user.ID, user.PasswordHash)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteUser(r.Context(), uid)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Poof, it's gone!"}); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createUserRequestToUser(req CreateUserRequest) swoleuser.User {
	return swoleuser.User{
		UserName:     req.UserName,
		PasswordHash: req.PasswordHash,
	}
}
