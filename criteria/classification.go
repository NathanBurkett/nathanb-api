package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldClassificationId = FieldId
const FieldClassificationTitle = "title"
const FieldClassificationSlug = "slug"
const FieldClassificationPublications = "publications"
const FieldClassificationCreatedAt = FieldCreatedAt
const FieldClassificationUpdatedAt = FieldUpdatedAt
const FieldClassificationDeletedAt = FieldDeletedAt

type SingleClassificationArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type classificationInterpretation struct{}

func (cl classificationInterpretation) fields() map[string]string {
	return map[string]string{
		FieldClassificationId:        data_object.FieldClassificationId,
		FieldClassificationTitle:     data_object.FieldClassificationTitle,
		FieldClassificationSlug:      data_object.FieldClassificationSlug,
		FieldClassificationCreatedAt: data_object.FieldClassificationCreatedAt,
		FieldClassificationUpdatedAt: data_object.FieldClassificationUpdatedAt,
		FieldClassificationDeletedAt: data_object.FieldClassificationDeletedAt,
	}
}

func (cl classificationInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SingleClassificationArgs:
		cl.interpretFirstClassificationArgs(c, T)
		break
	case PaginationArgs:
		T = cl.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown classification argument type: %s", reflect.TypeOf(T)))
	}
}

func (cl classificationInterpretation) interpretFirstClassificationArgs(c AbstractCriteria, args SingleClassificationArgs) {
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

func (cl classificationInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err    error
		skip   bool
	)

	column = cl.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown classification field: %s", field)
	}

	return column, skip, err
}

func (cl classificationInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldClassificationSlug, DirDesc),
		}
	}

	return args
}
