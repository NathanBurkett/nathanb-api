package data_object

import (
	"time"
	"github.com/satori/go.uuid"
)

const ProfileTable = "profile"

const FieldProfileId = FieldId
const FieldProfileFirstName = "first_name"
const FieldProfileLastName = "last_name"
const FieldProfileTwitterHandle = "twitter_handle"
const FieldProfileGithubHandle = "github_handle"
const FieldProfileCreatedAt = FieldCreatedAt
const FieldProfileUpdatedAt = FieldUpdatedAt
const FieldProfileDeletedAt = FieldDeletedAt

type Profile struct {
	ID            uuid.UUID `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	TwitterHandle string    `json:"twitterHandle"`
	GithubHandle  string    `json:"githubHandle"`
	UserID        uuid.UUID `json:"-"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
}

func (p Profile) Table() string {
	return ProfileTable
}