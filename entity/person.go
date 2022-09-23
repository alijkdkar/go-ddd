package entity

import (
	"github.com/google/uuid"
)

type person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
