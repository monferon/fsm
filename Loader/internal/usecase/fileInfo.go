package usecase

import (
	"context"
	"github.com/monferon/fsm/loader/internal/usecase/repo"
	"github.com/monferon/fsm/loader/internal/usecase/webapi"
)

type FileInfoUseCase struct {
	repo   FileInfoRepo
	webAPI FileInfoWebAPI
}

func New(r *repo.Repo, w *webapi.FileInfoWebAPI) *FileInfoUseCase {
	return &FileInfoUseCase{
		webAPI: w,
		repo:   r,
	}
}

func (fuc *FileInfoUseCase) Send(ctx context.Context) {

	//err := fuc.repo.Set(ctx)
	//if err != nil {
	//	return
	//}
	//fuc.webAPI.SendNext(ctx)
}
