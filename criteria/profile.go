package criteria

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	"fmt"
)

const FieldProfileId = FieldId
const FieldProfileFirstName = "firstName"
const FieldProfileLastName = "lastName"
const FieldProfileTwitterHandle = "twitterHandle"
const FieldProfileGithubHandle = "githubHandle"
const FieldProfileCreatedAt = FieldCreatedAt
const FieldProfileUpdatedAt = FieldUpdatedAt
const FieldProfileDeletedAt = FieldDeletedAt

type profileInterpretation struct {}

func (pi profileInterpretation) handleArgs(c *Criteria, args interface{}) {

}

func (pi profileInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err error
		skip bool
	)

	switch field {
	case FieldProfileId:
		column = data_object.FieldProfileId
		break
	case FieldProfileFirstName:
		column = data_object.FieldProfileFirstName
		break
	case FieldProfileLastName:
		column = data_object.FieldProfileLastName
		break
	case FieldProfileTwitterHandle:
		column = data_object.FieldProfileTwitterHandle
		break
	case FieldProfileGithubHandle:
		column = data_object.FieldProfileGithubHandle
		break
	case FieldProfileCreatedAt:
		column = data_object.FieldProfileCreatedAt
		break
	case FieldProfileUpdatedAt:
		column = data_object.FieldProfileUpdatedAt
		break
	case FieldProfileDeletedAt:
		column = data_object.FieldProfileDeletedAt
		break
	default:
		err = fmt.Errorf("unknown profile field: %s", field)
	}

	return column, skip, err
}
