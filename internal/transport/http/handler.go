package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type SwoleHandler struct {
	Router   *mux.Router
	UService UserService
	EService ExerciseService
	WService WorkoutService
	Server   *http.Server
}

func NewHandler(uservice UserService, eservice ExerciseService, wservice WorkoutService) *SwoleHandler {
	h := &SwoleHandler{
		UService: uservice,
		EService: eservice,
		WService: wservice,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: h.Router,
	}

	return h
}

func (h *SwoleHandler) mapRoutes() {
	h.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	}).Methods("GET")

	// User routes
	h.Router.HandleFunc("/api/v1/user", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", h.GetUserByID).Methods("GET")
	h.Router.HandleFunc("/api/v1/user/{username}", h.GetUserByName).Methods("GET")
	h.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", JWTAuth(h.UpdateUserPassword)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", JWTAuth(h.DeleteUser)).Methods("DELETE")
	h.Router.HandleFunc("/api/v1/auth", h.Login).Methods("POST")

	// Exercise routes
	h.Router.HandleFunc("/api/v1/exercise/name", h.GetExerciseByName).Methods("GET")
	h.Router.HandleFunc("/api/v1/exercise/{id:[0-9]+}", h.GetExerciseByID).Methods("GET")
	h.Router.HandleFunc("/api/v1/exercise/muscle", h.GetExerciseByMuscle).Methods("GET")
	h.Router.HandleFunc("/api/v1/exercise/equipment", h.GetExerciseByEquipment).Methods("GET")

	// Workout routes
	h.Router.HandleFunc("/api/v1/workout", JWTAuth(h.CreateWorkout)).Methods("POST")
	h.Router.HandleFunc("/api/v1/workout/user", JWTAuth(h.GetWorkoutsByUser)).Methods("GET")
	h.Router.HandleFunc("/api/v1/workout/{plan_id:[0-9]+}", JWTAuth(h.GetWorkoutExercises)).Methods("GET")
	h.Router.HandleFunc("/api/v1/workout", JWTAuth(h.UpdateWorkout)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/workout", JWTAuth(h.DeleteWorkout)).Methods("DELETE")
}

func (h *SwoleHandler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down")

	return nil
}
