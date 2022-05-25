package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type FileInfo struct {
	ID       uuid.UUID `gorm:"column:id"`
	Name     string    `gorm:"column:beast_id"`
	DtCreate time.Time `gorm:"column:dtCreate"`
	Path     string
	Author   string
	Size     int
}

func (FileInfo) TableName() string {
	return "files"
}
