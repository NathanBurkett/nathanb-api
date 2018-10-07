package data_object

import (
	"github.com/satori/go.uuid"
)

const TableCategoryPublication = "category_publication"

const FieldCategoryPublicationId = FieldId
const FieldCategoryPublicationCategoryId = "category_id"
const FieldCategoryPublicationPublicationId = "publication_id"

type CategoryPublication struct {
	ID            uuid.UUID `json:"-"`
	CategoryId    uuid.UUID `json:"-"`
	PublicationId uuid.UUID `json:"-"`
}

func (cp CategoryPublication) Table() string {
	return TableCategoryPublication
}

func (cp CategoryPublication) Fields() []string {
	return []string{
		FieldCategoryPublicationId,
		FieldCategoryPublicationCategoryId,
		FieldCategoryPublicationPublicationId,
	}
}
