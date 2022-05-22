package repo

import "github.com/monferon/fsm/loader/pkg/postgres"

type Repo struct {
	*postgres.Postgres
}

func New(*postgres.Postgres) *Repo {
	return &Repo{}
}
