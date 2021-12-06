package repos

import (
	project "github.com/PashaFedyushkin/GoProj"
	"github.com/jmoiron/sqlx"
)

type Authorizaton interface {
	CreateUser(user project.User) (int, error)
	GetUser(username, password string) (project.User, error)
}

type PlayGame interface {
}

type Statistic interface {
}

type Repos struct {
	Authorizaton
	PlayGame
	Statistic
}

func NewRepos(db *sqlx.DB) *Repos {
	return &Repos{
		Authorizaton: NewAuthPostgres(db),
	}
}
