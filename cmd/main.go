package main

import (
	"log"

	"github.com/Cr4z1k/MEDODS-test-task/internal/database"
	"github.com/Cr4z1k/MEDODS-test-task/internal/repository"
	"github.com/Cr4z1k/MEDODS-test-task/internal/service"
	"github.com/Cr4z1k/MEDODS-test-task/internal/transport/rest"
	"github.com/Cr4z1k/MEDODS-test-task/internal/transport/rest/handler"
	"github.com/Cr4z1k/MEDODS-test-task/pkg/auth"
)

func main() {
	s := new(rest.Server)

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal("Fatal creating db: ", err.Error())
	}

	tokenManager, err := auth.NewManager("salt123") //os.Getenv("tokenSalt"))
	if err != nil {
		log.Fatal("Fatal creating token manager: ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, tokenManager)
	handlers := handler.NewHandler(services)

	if err := s.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("Fatal starting server: ", err.Error())
	}
}
