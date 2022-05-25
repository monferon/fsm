package usecase

type FileInfoUseCase struct {
	webAPI fileInfoWebAPI
}

func New(w fileInfoWebAPI) *FileInfoUseCase {
	return &FileInfoUseCase{
		webAPI: w,
	}
}

func (fuc *FileInfoUseCase) Upload() {

}

func (fuc *FileInfoUseCase) SendEvent() {

}
