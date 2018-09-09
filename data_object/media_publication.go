package data_object

import "github.com/satori/go.uuid"

const TableMediaPublication = "media_publication"

type MediaPublication struct {
	MediaId    uuid.UUID `json:"-"`
	PublicationId uuid.UUID `json:"-"`
}

func (mp MediaPublication) Table() string {
	return TableMediaPublication
}
