package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

const FieldUserId = FieldId
const FieldUserEmail = "email"
const FieldUserPasswordDigest = "passwordDigest"
const FieldUserCreatedAt = FieldCreatedAt
const FieldUserUpdatedAt = FieldUpdatedAt
const FieldUserDeletedAt = FieldDeletedAt

type FirstUserArgs struct {
	ID    *uuid.UUID
	Email *string
}

type userInterpretation struct{}

func (ui userInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstUserArgs:
		ui.interpretFirstUserArgs(c, T)
		break
	case PaginationArgs:
		T = ui.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown category argument type: %s", T))
	}
}

func (ui userInterpretation) interpretFirstUserArgs(criteria AbstractCriteria, args FirstUserArgs) {
	if args.ID != nil {
		criteria.Where(query.Eq{data_object.FieldUserId: args.ID})
	}

	if args.Email != nil {
		criteria.Where(query.Eq{data_object.FieldUserEmail: args.Email})
	}
}

func (ui userInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err    error
		skip   bool
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

func (ui userInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	return args
}
