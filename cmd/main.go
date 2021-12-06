package main

import (
	"log"

	_ "github.com/lib/pq"

	pb "github.com/PashaFedyushkin/GoProj"
	"github.com/PashaFedyushkin/GoProj/pkg/handler"
	"github.com/PashaFedyushkin/GoProj/pkg/repos"
	"github.com/PashaFedyushkin/GoProj/pkg/service"
)

func main() {
	db, err := repos.NewPostgresDB(repos.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Failed initialized DB: %s", err.Error())
	}

	rep := repos.NewRepos(db)
	services := service.NewService(*rep)
	handlers := handler.NewHandler(services)

	srv := new(pb.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
