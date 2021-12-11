package main

import (
	"log"
	"net"

	"github.com/PashaFedyushkin/GoProj/pkg/api"

	"github.com/PashaFedyushkin/GoProj/pkg/service"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &service.GAMEServer{}
	api.RegisterGameServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
