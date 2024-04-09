package main

import (
	"log"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/handler"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/repository"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(banner.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
