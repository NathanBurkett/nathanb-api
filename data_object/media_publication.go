package data_object

import "github.com/satori/go.uuid"

const TableMediaPublication = "media_publication"

const FieldMediaPublicationId = FieldId
const FieldMediaPublicationMediaId = "media_id"
const FieldMediaPublicationPublicationId = "publication_id"

type MediaPublication struct {
	ID            uuid.UUID `json:"-"`
	MediaId       uuid.UUID `json:"-"`
	PublicationId uuid.UUID `json:"-"`
}

func (mp MediaPublication) Table() string {
	return TableMediaPublication
}

func (mp MediaPublication) Fields() []string {
	return []string{
		FieldMediaPublicationId,
		FieldMediaPublicationMediaId,
		FieldMediaPublicationPublicationId,
	}
}
