package data_object

import (
	"github.com/satori/go.uuid"
	"time"
)

const TableCategory = "category"

const FieldCategoryId = FieldId
const FieldCategoryTitle = "title"
const FieldCategorySlug = "slug"
const FieldCategoryCreatedAt = FieldCreatedAt
const FieldCategoryUpdatedAt = FieldUpdatedAt
const FieldCategoryDeletedAt = FieldDeletedAt

type Category struct {
	ID           uuid.UUID `db:"id"`
	Title        string    `db:"title"`
	Slug         string    `db:"slug"`
	Publications []*Publication
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DeletedAt    time.Time `db:"deleted_at"`
}

func (c Category) Table() string {
	return TableCategory
}

func (c Category) Fields() []string {
	return []string{
		FieldCategoryId,
		FieldCategoryTitle,
		FieldCategorySlug,
		FieldCategoryCreatedAt,
		FieldCategoryUpdatedAt,
		FieldCategoryDeletedAt,
	}
}
