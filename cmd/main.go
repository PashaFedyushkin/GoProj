package main

import (
	"log"

	pb "github.com/PashaFedyushkin/GoProj"
)

func main() {
	srv := new(pb.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
