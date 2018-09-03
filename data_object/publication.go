package data_object

import (
	"time"
	"github.com/satori/go.uuid"
)

const TablePublication = "publication"

const FieldPublicationId = FieldId
const FieldPublicationTitle = "title"
const FieldPublicationSlug = "slug"
const FieldPublicationPublishedAt = "published_at"
const FieldPublicationCreatedAt = FieldCreatedAt
const FieldPublicationUpdatedAt = FieldUpdatedAt
const FieldPublicationDeletedAt = FieldDeletedAt

type Publication struct {
	ID               uuid.UUID      `json:"id"`
	Title            string         `json:"title"`
	Slug             string         `json:"slug"`
	ClassificationID uuid.UUID      `json:"-"`
	Classification   *Classification `json:"classification"`
	ContentBlocks    []*ContentBlock `json:"contentBlocks"`
	Categories       []*Category     `json:"categories"`
	Media            []*Media        `json:"media"`
	Users            []*User         `json:"users"`
	PublishedAt      time.Time      `json:"publishedAt"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        time.Time      `json:"deletedAt"`
}

func (p Publication) Table() string {
	return TablePublication
}
