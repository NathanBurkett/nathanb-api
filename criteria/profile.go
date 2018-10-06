package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

const FieldProfileId = FieldId
const FieldProfileFirstName = "firstName"
const FieldProfileLastName = "lastName"
const FieldProfileTwitterHandle = "twitterHandle"
const FieldProfileGithubHandle = "githubHandle"
const FieldProfileCreatedAt = FieldCreatedAt
const FieldProfileUpdatedAt = FieldUpdatedAt
const FieldProfileDeletedAt = FieldDeletedAt

type FirstProfileArgs struct {
	ID            *uuid.UUID
	UserID        *string
	GithubHandle  *string
	TwitterHandle *string
}

type profileInterpretation struct{}

func (pi profileInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstProfileArgs:
		pi.interpretFirstCategoryArgs(c, T)
		break
	case PaginationArgs:
		T = pi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown category argument type: %s", T))
	}
}

func (pi profileInterpretation) interpretFirstCategoryArgs(criteria AbstractCriteria, args FirstProfileArgs) {
	if args.ID != nil {
		criteria.Where(query.Eq{data_object.FieldProfileId: args.ID})
	}

	if args.UserID != nil {
		criteria.Where(query.Eq{data_object.FieldProfileUserId: args.UserID})
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

func (pi profileInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy != nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldProfileLastName, DirDesc),
			fmt.Sprintf("%s %s", data_object.FieldProfileFirstName, DirDesc),
		}
	}

	return args
}
