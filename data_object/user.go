package data_object

import (
	"time"
	"github.com/satori/go.uuid"
)

const TableUser = "user"

const FieldUserId = FieldId
const FieldUserEmail = "email"
const FieldUserPasswordDigest = "password_digest"
const FieldUserCreatedAt = FieldCreatedAt
const FieldUserUpdatedAt = FieldUpdatedAt
const FieldUserDeletedAt = FieldDeletedAt

type User struct {
	ID             uuid.UUID     `json:"id"`
	Email          string        `json:"email"`
	PasswordDigest string        `json:"-"`
	Profile        Profile       `json:"profile"`
	Publications   []Publication `json:"publications"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	DeletedAt      time.Time     `json:"deletedAt"`
}

func (u User) Table() string {
	return TableUser
}
