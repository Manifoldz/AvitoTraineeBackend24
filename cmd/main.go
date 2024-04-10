package main

import (
	"log"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/handler"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/repository"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		Password: viper.GetString("db.Password"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(banner.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
