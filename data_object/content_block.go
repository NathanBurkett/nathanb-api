package data_object

import (
	"github.com/satori/go.uuid"
	"time"
)

const TableContentBlock = "content_block"

const FieldContentBlockId = FieldId
const FieldContentBlockType = "type"
const FieldContentBlockContent = "content"
const FieldContentBlockOrder = "order"
const FieldContentBlockCreatedAt = FieldCreatedAt
const FieldContentBlockUpdatedAt = FieldUpdatedAt
const FieldContentBlockDeletedAt = FieldDeletedAt

type ContentBlock struct {
	ID            uuid.UUID   `json:"id"`
	Type          string      `json:"type"`
	Content       string      `json:"content"`
	Order         uint8       `json:"order"`
	Publication   Publication `json:"publication"`
	PublicationId uuid.UUID   `json:"-"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	DeletedAt     time.Time   `json:"deletedAt"`
}

func (cb ContentBlock) Table() string {
	return TableContentBlock
}

func (cb ContentBlock) Fields() []string {
	return []string{
		FieldContentBlockId,
		FieldContentBlockType,
		FieldContentBlockContent,
		FieldContentBlockOrder,
		FieldContentBlockCreatedAt,
		FieldContentBlockUpdatedAt,
		FieldContentBlockDeletedAt,
	}
}
