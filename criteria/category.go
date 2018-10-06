package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

const FieldCategoryId = FieldId
const FieldCategoryTitle = "title"
const FieldCategorySlug = "slug"
const FieldCategoryPublications = "publications"
const FieldCategoryCreatedAt = FieldCreatedAt
const FieldCategoryUpdatedAt = FieldUpdatedAt
const FieldCategoryDeletedAt = FieldDeletedAt

type FirstCategoryArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type categoryInterpretation struct {}

func (ci categoryInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstCategoryArgs:
		ci.interpretFirstCategoryArgs(c, T)
		break
	case PaginationArgs:
		T = ci.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown category argument type: %s", T))
	}
}

func (ci categoryInterpretation) interpretFirstCategoryArgs(c AbstractCriteria, args FirstCategoryArgs) {
	if args.ID != nil {
		c.Where(query.Eq{data_object.FieldCategoryId: args.ID})
	}

	if args.Title != nil {
		c.Where(query.Eq{data_object.FieldCategoryTitle: args.Title})
	}

	if args.Slug != nil {
		c.Where(query.Eq{data_object.FieldCategorySlug: args.Slug})
	}
}

func (ci categoryInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err error
		skip bool
	)

	switch field {
	case FieldCategoryId:
		column = data_object.FieldCategoryId
	case FieldCategoryTitle:
		column = data_object.FieldCategoryTitle
	case FieldCategorySlug:
		column = data_object.FieldCategorySlug
	case FieldCategoryPublications:
		skip = true
	case FieldCategoryCreatedAt:
		column = data_object.FieldCategoryCreatedAt
	case FieldCategoryUpdatedAt:
		column = data_object.FieldCategoryUpdatedAt
	case FieldCategoryDeletedAt:
		column = data_object.FieldCategoryDeletedAt
	default:
		err = fmt.Errorf("unknown category field: %s", field)
	}

	return column, skip, err
}

func (ci categoryInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy != nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldCategorySlug, DirDesc),
		}
	}

	return args
}
