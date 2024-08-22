package startFunc

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/benodiwal/gorm-studio/pkg/database"
	"github.com/benodiwal/gorm-studio/pkg/env"
	"github.com/benodiwal/gorm-studio/pkg/routes"
)

func Start() {
	env.Load()
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		runClient()
	}()

	go func() {
		defer wg.Done()
		runServer()
	}()

	wg.Wait()
}

func runServer() {
	logger := log.Default()
	db := database.Connect(logger)

	router := routes.New(db)

	router.RegisterMiddlewares()
	router.RegisterRoutes()

	router.Engine.Run(":8000")
}

func runClient() {
	clientDir := filepath.Join("..", "client")
	err := os.Chdir(clientDir)
	if err != nil {
		fmt.Println("Failed to change directory: ", err)
		return
	}

	command := exec.Command("yarn", "dev")
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
		fmt.Println("Failed to start the dev server: ", err)
	}
}
