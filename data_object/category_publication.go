package data_object

import (
	"github.com/satori/go.uuid"
)

const CategoryPublicationTable = "category_publication"

type CategoryPublication struct {
	CategoryId    uuid.UUID `json:"-"`
	PublicationId uuid.UUID `json:"-"`
}

func (cp CategoryPublication) Table() string {
	return CategoryPublicationTable
}
