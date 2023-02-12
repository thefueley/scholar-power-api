package main

import (
	"fmt"

	"github.com/thefueley/scholar-power-api/internal/db"
	"github.com/thefueley/scholar-power-api/internal/exercise"
	"github.com/thefueley/scholar-power-api/internal/history"
	transportHttp "github.com/thefueley/scholar-power-api/internal/transport/http"
	swoleuser "github.com/thefueley/scholar-power-api/internal/user"
	"github.com/thefueley/scholar-power-api/internal/workout"

	log "github.com/sirupsen/logrus"
)

func Run() error {
	fmt.Println("Starting Scholar Power API server")

	store, err := db.NewDatabase("internal/db/scholarpower.db")
	if err != nil {
		fmt.Println("error opening database")
		return err
	}

	if err := store.Migrate(); err != nil {
		fmt.Println("error migrating database")
		return err
	}

	userService := swoleuser.NewUserService(store)
	exerciseService := exercise.NewExerciseService(store)
	workoutService := workout.NewWorkoutService(store)
	historyService := history.NewHistoryService(store)

	httpHandler := transportHttp.NewHandler(userService, exerciseService, workoutService, historyService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting Scholar Power API")
	}
}
