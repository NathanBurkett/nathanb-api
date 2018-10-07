package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldProfileId = FieldId
const FieldProfileUserId = "userId"
const FieldProfileFirstName = "firstName"
const FieldProfileLastName = "lastName"
const FieldProfileTwitterHandle = "twitterHandle"
const FieldProfileGithubHandle = "githubHandle"
const FieldProfileCreatedAt = FieldCreatedAt
const FieldProfileUpdatedAt = FieldUpdatedAt
const FieldProfileDeletedAt = FieldDeletedAt

type SingleProfileArgs struct {
	ID            *uuid.UUID
	UserId        *uint64
	GithubHandle  *string
	TwitterHandle *string
}

type profileInterpretation struct{}

func (profileInterpretation) fields() map[string]string {
	return map[string]string{
		FieldProfileId:            data_object.FieldProfileId,
		FieldProfileFirstName:     data_object.FieldProfileFirstName,
		FieldProfileLastName:      data_object.FieldProfileLastName,
		FieldProfileTwitterHandle: data_object.FieldProfileTwitterHandle,
		FieldProfileGithubHandle:  data_object.FieldProfileGithubHandle,
		FieldProfileUserId:        data_object.FieldProfileUserId,
		FieldProfileCreatedAt:     data_object.FieldProfileCreatedAt,
		FieldProfileUpdatedAt:     data_object.FieldProfileUpdatedAt,
		FieldProfileDeletedAt:     data_object.FieldProfileDeletedAt,
	}
}

func (pi profileInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SingleProfileArgs:
		pi.interpretFirstCategoryArgs(c, T)
		break
	case PaginationArgs:
		T = pi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown profile argument type: %s", reflect.TypeOf(T)))
	}
}

func (pi profileInterpretation) interpretFirstCategoryArgs(criteria AbstractCriteria, args SingleProfileArgs) {
	if args.ID != nil {
		criteria.Where(query.Eq{data_object.FieldProfileId: args.ID})
	}

	if args.UserId != nil {
		criteria.Where(query.Eq{data_object.FieldProfileUserId: args.UserId})
	}

	if args.GithubHandle != nil {
		criteria.Where(query.Eq{data_object.FieldProfileGithubHandle: args.GithubHandle})
	}

	if args.TwitterHandle != nil {
		criteria.Where(query.Eq{data_object.FieldProfileTwitterHandle: args.TwitterHandle})
	}
}

func (pi profileInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err    error
		skip   bool
	)

	column = pi.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown field: %s", field)
	}

	return column, skip, err
}

func (pi profileInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldProfileLastName, DirDesc),
			fmt.Sprintf("%s %s", data_object.FieldProfileFirstName, DirDesc),
		}
	}

	return args
}
