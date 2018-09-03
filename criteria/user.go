package criteria

import (
	"fmt"
	"github.com/nathanburkett/nathanb-api/data_object"
)

const FieldUserId = FieldId
const FieldUserEmail = "email"
const FieldUserPasswordDigest = "passwordDigest"
const FieldUserCreatedAt = FieldCreatedAt
const FieldUserUpdatedAt = FieldUpdatedAt
const FieldUserDeletedAt = FieldDeletedAt

type userInterpretation struct {}

func (ui userInterpretation) handleArgs(c *Criteria, args interface{}) {

}

func (ui userInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err error
		skip bool
	)

	switch field {
	case FieldUserId:
		column = data_object.FieldUserId
		break
	case FieldUserEmail:
		column = data_object.FieldUserEmail
		break
	case FieldUserPasswordDigest:
		column = data_object.FieldUserPasswordDigest
		break
	case FieldUserCreatedAt:
		column = data_object.FieldUserCreatedAt
		break
	case FieldUserUpdatedAt:
		column = data_object.FieldUserUpdatedAt
		break
	case FieldUserDeletedAt:
		column = data_object.FieldUserDeletedAt
		break
	default:
		err = fmt.Errorf("unknown user field: %s", field)
	}

	return column, skip, err
}
