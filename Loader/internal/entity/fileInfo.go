package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type FileInfo struct {
	ID       uuid.UUID
	Name     string
	DtCreate time.Time
	Path     string
	Author   string
	Size     int
}
