package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	swoleuser "github.com/thefueley/scholar-power-api/internal/user"
)

type UserService interface {
	CreateUser(ctx context.Context, username string, password string) (string, error)
	GetUserByID(ctx context.Context, uid string) (swoleuser.User, error)
	GetUserByName(ctx context.Context, username string) (swoleuser.User, error)
	UpdateUserPassword(ctx context.Context, uid string, password string) error
	DeleteUser(ctx context.Context, uid string) error
	Login(ctx context.Context, username string, password string) (string, error)
}

type Response struct {
	Message string
}

type CreateUserRequest struct {
	UserName     string `json:"username" validate:"required"`
	PasswordHash string `json:"password" validate:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *SwoleHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	newUser, err := h.UService.CreateUser(r.Context(), convertedUser.UserName, convertedUser.PasswordHash)
	if err != nil {
		http.Error(w, "user already exists", http.StatusForbidden)
		return
	}

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := h.UService.GetUserByID(r.Context(), uid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	username := usr.UserName

	if err := json.NewEncoder(w).Encode(username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uname := vars["username"]
	if uname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := h.UService.GetUserByName(r.Context(), uname)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uid := usr.ID

	if err := json.NewEncoder(w).Encode(uid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SwoleHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	if uid == "" {
		fmt.Println("view.UpdateUserPassword: uid is empty")
		w.WriteHeader(http.StatusBadRequest)
	}

	var newPassword LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&newPassword); err != nil {
		fmt.Printf("view.UpdateUserPassword NewDecoder: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err := h.AuthZ(r, uid)
	if err != nil {
		fmt.Printf("view.UpdateUserPassword AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.UService.UpdateUserPassword(r.Context(), uid, newPassword.Password)
	if err != nil {
		fmt.Printf("view.UpdateUserPassword: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode("password updated"); err != nil {
		fmt.Printf("view.UpdateUserPassword NewEncoder: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *SwoleHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.AuthZ(r, uid)
	if err != nil {
		fmt.Printf("view.DeletePassword AuthZ: %v\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.UService.DeleteUser(r.Context(), uid)
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

func (h *SwoleHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		fmt.Printf("view.Login NewDecoder: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.UService.Login(r.Context(), userLogin.UserName, userLogin.Password)
	if err != nil {
		fmt.Printf("view.Login Login: %v\n", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := json.NewEncoder(w).Encode(token); err != nil {
		fmt.Printf("view.Login NewEncoder: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createUserRequestToUser(req CreateUserRequest) swoleuser.User {
	return swoleuser.User{
		UserName:     req.UserName,
		PasswordHash: req.PasswordHash,
	}
}

func (h *SwoleHandler) AuthZ(r *http.Request, OwnerID string) error {
	authHeader := r.Header["Authorization"]
	authHeaderParts := strings.Split(authHeader[0], " ")
	requestor, err := Requestor(authHeaderParts[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	requestorID, err := h.UService.GetUserByID(context.Background(), requestor)

	if err != nil {
		fmt.Printf("view.user: AuthZ: %v\n", err)
		return err
	}

	if requestorID.ID != OwnerID {
		fmt.Println("view.AuthZ: requestorID.ID != OwnerID")
		return errors.New("Unauthorized")
	}

	return nil
}
