package entity

import "github.com/google/uuid"

type item struct {
	ID          uuid.UUID
	name        string
	Description string
}
