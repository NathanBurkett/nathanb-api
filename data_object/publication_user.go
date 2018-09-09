package data_object

import "github.com/satori/go.uuid"

const TablePublicationUser = "publication_user"

type PublicationUser struct {
	PublicationId uuid.UUID `json:"-"`
	UserId        uuid.UUID `json:"-"`
}

func (pu PublicationUser) Table() string {
	return TablePublicationUser
}
