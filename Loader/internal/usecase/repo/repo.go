package repo

import (
	"context"
	"github.com/monferon/fsm/loader/internal/entity"
	"github.com/monferon/fsm/loader/pkg/postgres"
)

type Repo struct {
	*postgres.Postgres
}

func New(*postgres.Postgres) *Repo {
	return &Repo{}
}

func (r *Repo) Set(ctx context.Context, info entity.FileInfo) error {
	//INSERT INTO TABLE SET ()
	//TODO Add set to DB (gorm)
	//fl := entity.FileInfo{}

	return nil
}
func (r *Repo) Get(ctx context.Context, filename string) (entity.FileInfo, error) {
	//GET WHERE DRATATA
	//TODO Add get from DB (gorm)
	return entity.FileInfo{}, nil
}
