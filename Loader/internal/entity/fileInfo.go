package entity

import "github.com/gofrs/uuid"

type FileInfo struct {
	ID     uuid.UUID
	Name   string
	Author string
	Size   int
}
