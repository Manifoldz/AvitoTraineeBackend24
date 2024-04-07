package main

import (
	"log"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
)

func main() {
	srv := new(banner.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
