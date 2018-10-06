package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

const FieldClassificationId = FieldId
const FieldClassificationTitle = "title"
const FieldClassificationSlug = "slug"
const FieldClassificationPublications = "publications"
const FieldClassificationCreatedAt = FieldCreatedAt
const FieldClassificationUpdatedAt = FieldUpdatedAt
const FieldClassificationDeletedAt = FieldDeletedAt

type FirstClassificationArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type classificationInterpretation struct{}

func (cl classificationInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstClassificationArgs:
		cl.interpretFirstClassificationArgs(c, T)
		break
	case PaginationArgs:
		T = cl.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown classification argument type: %s", T))
	}
}

func (cl classificationInterpretation) interpretFirstClassificationArgs(c AbstractCriteria, args FirstClassificationArgs) {
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
	)

	shouldSkip := false

	switch field {
	case FieldClassificationId:
		column = data_object.FieldClassificationId
		break
	case FieldClassificationTitle:
		column = data_object.FieldClassificationTitle
		break
	case FieldClassificationSlug:
		column = data_object.FieldClassificationSlug
		break
	case FieldClassificationPublications:
		shouldSkip = true
		break
	case FieldClassificationCreatedAt:
		column = data_object.FieldClassificationCreatedAt
		break
	case FieldClassificationUpdatedAt:
		column = data_object.FieldClassificationUpdatedAt
		break
	case FieldClassificationDeletedAt:
		column = data_object.FieldClassificationDeletedAt
		break
	default:
		err = fmt.Errorf("unknown classification field: %s", field)
	}

	return column, shouldSkip, err
}

func (cl classificationInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy != nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldClassificationSlug, DirDesc),
		}
	}

	return args
}
