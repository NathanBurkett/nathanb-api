package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldCategoryId = FieldId
const FieldCategoryTitle = "title"
const FieldCategorySlug = "slug"
const FieldCategoryPublications = "publications"
const FieldCategoryCreatedAt = FieldCreatedAt
const FieldCategoryUpdatedAt = FieldUpdatedAt
const FieldCategoryDeletedAt = FieldDeletedAt

type SingleCategoryArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type categoryInterpretation struct{}

func (ci categoryInterpretation) fields() map[string]string {
	return map[string]string{
		FieldCategoryId:        data_object.FieldCategoryId,
		FieldCategoryTitle:     data_object.FieldCategoryTitle,
		FieldCategorySlug:      data_object.FieldCategorySlug,
		FieldCategoryCreatedAt: data_object.FieldCategoryCreatedAt,
		FieldCategoryUpdatedAt: data_object.FieldCategoryUpdatedAt,
		FieldCategoryDeletedAt: data_object.FieldCategoryDeletedAt,
	}
}

func (ci categoryInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SingleCategoryArgs:
		ci.interpretFirstCategoryArgs(c, T)
		break
	case PaginationArgs:
		T = ci.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown category argument type: %s", reflect.TypeOf(T)))
	}
}

func (ci categoryInterpretation) interpretFirstCategoryArgs(c AbstractCriteria, args SingleCategoryArgs) {
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
		err    error
		skip   bool
	)

	column = ci.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown field: %s", field)
	}

	return column, skip, err
}

func (ci categoryInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldCategorySlug, DirDesc),
		}
	}

	return args
}
