package webapi

import (
	"context"
	"github.com/monferon/fsm/loader/internal/entity"
)

type FileInfoWebAPI struct {
}

func New() *FileInfoWebAPI {
	return &FileInfoWebAPI{}
}

func (w *FileInfoWebAPI) SendNext(ctx context.Context, info entity.FileInfo) error {
	return nil
}
