package usecase

type (
	fileInfo interface {
		Upload()
		SendEvent()
	}
	fileInfoWebAPI interface {
		SendEvent() error
	}
)
