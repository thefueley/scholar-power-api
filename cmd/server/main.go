package main

import (
	"fmt"

	"github.com/thefueley/scholar-power-api/internal/db"
	transportHttp "github.com/thefueley/scholar-power-api/internal/transport/http"
	swoleuser "github.com/thefueley/scholar-power-api/internal/user"

	log "github.com/sirupsen/logrus"
)

func Run() error {
	fmt.Println("Starting API server")

	store, err := db.NewDatabase()
	if err != nil {
		fmt.Println("error opening database")
		return err
	}

	if err := store.Migrate(); err != nil {
		fmt.Println("error migrating database")
		return err
	}

	userService := swoleuser.NewUserService(store)

	httpHandler := transportHttp.NewHandler(userService)
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
