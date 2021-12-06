package service

import (
	project "github.com/PashaFedyushkin/GoProj"
	"github.com/PashaFedyushkin/GoProj/pkg/repos"
	_ "github.com/lib/pq"
)

type Authorizaton interface {
	CreateUser(user project.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type PlayGame interface {
}

type Statistic interface {
}

type Service struct {
	Authorizaton
	PlayGame
	Statistic
}

func NewService(rep repos.Repos) *Service {
	return &Service{
		Authorizaton: NewAuthService(rep.Authorizaton),
	}
}
