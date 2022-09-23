package aggregate

import (
	"github.com/alijkdkar/go-ddd/valueobjects"

	"github.com/alijkdkar/go-ddd/entity"
)

type Customer struct {
	person      *entity.Person
	product     []*entity.Item
	transaction []*valueobjects.Transaction
}
