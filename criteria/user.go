package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldUserId = FieldId
const FieldUserEmail = "email"
const FieldUserPasswordDigest = "passwordDigest"
const FieldUserCreatedAt = FieldCreatedAt
const FieldUserUpdatedAt = FieldUpdatedAt
const FieldUserDeletedAt = FieldDeletedAt

type SingleUserArgs struct {
	ID    *uuid.UUID
	Email *string
}

type userInterpretation struct{}

func (userInterpretation) fields() map[string]string {
	return map[string]string{
		FieldUserId:             data_object.FieldUserId,
		FieldUserEmail:          data_object.FieldUserEmail,
		FieldUserPasswordDigest: data_object.FieldUserPasswordDigest,
		FieldUserCreatedAt:      data_object.FieldUserCreatedAt,
		FieldUserUpdatedAt:      data_object.FieldUserUpdatedAt,
		FieldUserDeletedAt:      data_object.FieldUserDeletedAt,
	}
}

func (ui userInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SingleUserArgs:
		ui.interpretFirstUserArgs(c, T)
		break
	case PaginationArgs:
		T = ui.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown user argument type: %s", reflect.TypeOf(T)))
	}
}

func (ui userInterpretation) interpretFirstUserArgs(criteria AbstractCriteria, args SingleUserArgs) {
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

	column = ui.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown field: %s", field)
	}

	return column, skip, err
}

func (ui userInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{}
	}

	return args
}
