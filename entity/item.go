package entity

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID
	name        string
	Description string
}
