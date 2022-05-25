package webapi

type FileInfoWebAPI struct {
}

func New() *FileInfoWebAPI {
	return &FileInfoWebAPI{}
}

func (fw *FileInfoWebAPI) SendEvent() error {
	return nil
}
