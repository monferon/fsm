package repo

import (
	"context"
	"fmt"
	"github.com/monferon/fsm/loader/internal/entity"
	"github.com/monferon/fsm/loader/pkg/postgres"
)

type Repo struct {
	*postgres.Postgres
}

func New(*postgres.Postgres) *Repo {
	return &Repo{}
}

func (r *Repo) Set(ctx context.Context, file entity.FileInfo) error {
	//INSERT INTO TABLE SET ()
	//TODO test this
	result := r.DB.Create(file)
	fmt.Println(result)

	//fl := entity.FileInfo{}

	return nil
}
func (r *Repo) Get(ctx context.Context, filename string) (entity.FileInfo, error) {
	//GET WHERE DRATATA
	//TODO test this
	file := entity.FileInfo{}
	r.DB.First(&file, "Name = ?", filename)
	return entity.FileInfo{}, nil
}
