package server

import (
	"context"

	"github.com/PashaFedyushkin/GoProj/pkg/api"
)

type AUTHServer struct{}

func (s *AUTHServer) Auth(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	var a string
	a = "123"
	var b string
	b = "123"
	if req.GetLogin() == a && req.GetPassword() == b {
		return &api.AuthResponse{Res: "Hello"}, nil
	}
	return &api.AuthResponse{Res: "Error"}, nil
}
