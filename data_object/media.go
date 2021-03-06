package data_object

import (
	"github.com/satori/go.uuid"
	"time"
)

const TableMedia = "media"

const FieldMediaId = FieldId
const FieldMediaType = "type"
const FieldMediaSubtype = "subtype"
const FieldMediaTitle = "title"
const FieldMediaPath = "path"
const FieldMediaAlt = "alt"
const FieldMediaAspectRatio = "aspect_ratio"
const FieldMediaCreatedAt = FieldCreatedAt
const FieldMediaUpdatedAt = FieldUpdatedAt
const FieldMediaDeletedAt = FieldDeletedAt

type Media struct {
	ID           uuid.UUID      `json:"id"`
	Type         string         `json:"type"`
	Subtype      string         `json:"subtype"`
	Title        string         `json:"title"`
	Path         string         `json:"path"`
	Alt          string         `json:"alt"`
	AspectRatio  string         `json:"aspectRatio"`
	Publications []*Publication `json:"publications"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    time.Time      `json:"deletedAt"`
}

func (m Media) Table() string {
	return TableMedia
}

func (m Media) Fields() []string {
	return []string{
		FieldMediaId,
		FieldMediaType,
		FieldMediaSubtype,
		FieldMediaTitle,
		FieldMediaPath,
		FieldMediaAlt,
		FieldMediaAspectRatio,
		FieldMediaCreatedAt,
		FieldMediaUpdatedAt,
		FieldMediaDeletedAt,
	}
}
