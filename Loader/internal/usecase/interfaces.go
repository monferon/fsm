package usecase

import (
	"context"
	"github.com/monferon/fsm/loader/internal/entity"
)

type (
	FileInfo interface {
		Send(ctx context.Context)
	}

	FileInfoRepo interface {
		Set(ctx context.Context, info entity.FileInfo) error
		Get(ctx context.Context, filename string) (entity.FileInfo, error)
	}
	FileInfoWebAPI interface {
		SendNext(ctx context.Context, info entity.FileInfo) error
	}
)
