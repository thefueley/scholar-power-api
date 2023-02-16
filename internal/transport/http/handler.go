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
	HService HistoryService
	Server   *http.Server
}

func NewHandler(uservice UserService, eservice ExerciseService, wservice WorkoutService, hservice HistoryService) *SwoleHandler {
	h := &SwoleHandler{
		UService: uservice,
		EService: eservice,
		WService: wservice,
		HService: hservice,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(mux.CORSMethodMiddleware(h.Router))
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
	}).Methods(http.MethodGet, http.MethodOptions)

	// User routes
	h.Router.HandleFunc("/api/v1/register", h.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", h.GetUserByID).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/user/{username}", h.GetUserByName).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", JWTAuth(h.UpdateUserPassword)).Methods(http.MethodPut, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", JWTAuth(h.DeleteUser)).Methods(http.MethodDelete, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/auth", h.Login).Methods(http.MethodPost, http.MethodOptions)

	// Exercise routes
	h.Router.HandleFunc("/api/v1/exercise/name", h.GetExerciseByName).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/exercise/{id:[0-9]+}", h.GetExerciseByID).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/exercise/{muscle}", h.GetExerciseByMuscle).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/exercise/equipment", h.GetExerciseByEquipment).Methods(http.MethodGet, http.MethodOptions)

	// Workout routes
	h.Router.HandleFunc("/api/v1/workout", JWTAuth(h.CreateWorkout)).Methods(http.MethodPost, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/workout/user/{username}", JWTAuth(h.GetWorkoutsByUser)).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/workout/{plan_id:[0-9]+}", JWTAuth(h.GetWorkoutExercises)).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/workout", JWTAuth(h.UpdateWorkout)).Methods(http.MethodPut, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/workout", JWTAuth(h.DeleteWorkout)).Methods(http.MethodDelete, http.MethodOptions)

	// History routes
	h.Router.HandleFunc("/api/v1/history", JWTAuth(h.CreateHistory)).Methods(http.MethodPost, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/history/{id:[0-9]+}", JWTAuth(h.GetHistory)).Methods(http.MethodGet, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/history", JWTAuth(h.UpdateHistory)).Methods(http.MethodPut, http.MethodOptions)
	h.Router.HandleFunc("/api/v1/history", JWTAuth(h.DeleteHistory)).Methods(http.MethodDelete, http.MethodOptions)

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
