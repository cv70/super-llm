package gid

import "github.com/google/uuid"

func NewUUID() (uuid.UUID, error) {
	return uuid.NewV7()
}
