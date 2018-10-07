package data_object

import (
	"github.com/satori/go.uuid"
	"time"
)

const TableClassification = "classification"

const FieldClassificationId = FieldId
const FieldClassificationTitle = "title"
const FieldClassificationSlug = "slug"
const FieldClassificationCreatedAt = FieldCreatedAt
const FieldClassificationUpdatedAt = FieldUpdatedAt
const FieldClassificationDeletedAt = FieldDeletedAt

type Classification struct {
	ID           uuid.UUID      `json:"id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug"`
	Publications []*Publication `json:"publications"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    time.Time      `json:"deletedAt"`
}

func (c Classification) Table() string {
	return TableClassification
}

func (c Classification) Fields() []string {
	return []string{
		FieldClassificationId,
		FieldClassificationTitle,
		FieldClassificationSlug,
		FieldClassificationCreatedAt,
		FieldClassificationUpdatedAt,
		FieldClassificationDeletedAt,
	}
}
