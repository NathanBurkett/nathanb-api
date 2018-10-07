package data_object

import "github.com/satori/go.uuid"

const TablePublicationUser = "publication_user"

const FieldPublicationUserId = FieldId
const FieldPublicationUserPublicationId = "publication_id"
const FieldPublicationUserUserId = "user_id"

type PublicationUser struct {
	ID            uuid.UUID `json:"-"`
	PublicationId uuid.UUID `json:"-"`
	UserId        uuid.UUID `json:"-"`
}

func (pu PublicationUser) Table() string {
	return TablePublicationUser
}

func (pu PublicationUser) Fields() []string {
	return []string{
		FieldPublicationUserId,
		FieldPublicationUserPublicationId,
		FieldPublicationUserUserId,
	}
}
