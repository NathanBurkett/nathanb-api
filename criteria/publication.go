package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

const FieldPublicationId = FieldId
const FieldPublicationTitle = "title"
const FieldPublicationSlug = "slug"
const FieldPublicationPublishedAt = "publishedAt"
const FieldPublicationCreatedAt = FieldCreatedAt
const FieldPublicationUpdatedAt = FieldUpdatedAt
const FieldPublicationDeletedAt = FieldDeletedAt

type FirstPublicationArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type publicationInterpretation struct {}

func (pi publicationInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstPublicationArgs:
		pi.interpretFirstPublicationArgs(c, T)
		break
	case PaginationArgs:
		T = pi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown category argument type: %s", T))
	}
}

func (pi publicationInterpretation) interpretFirstPublicationArgs(criteria AbstractCriteria, args FirstPublicationArgs) {
	if args.ID != nil {
		criteria.Where(query.Eq{data_object.FieldPublicationId: args.ID})
	}

	if args.Title != nil {
		criteria.Where(query.Eq{data_object.FieldPublicationTitle: args.Title})
	}

	if args.Slug != nil {
		criteria.Where(query.Eq{data_object.FieldPublicationSlug: args.Slug})
	}
}

func (pi publicationInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err error
		skip bool
	)

	switch field {
	case FieldPublicationId:
		column = data_object.FieldPublicationId
		break
	case FieldPublicationTitle:
		column = data_object.FieldPublicationTitle
		break
	case FieldPublicationSlug:
		column = data_object.FieldPublicationSlug
		break
	case FieldPublicationPublishedAt:
		column = data_object.FieldPublicationPublishedAt
		break
	case FieldPublicationCreatedAt:
		column = data_object.FieldPublicationCreatedAt
		break
	case FieldPublicationUpdatedAt:
		column = data_object.FieldPublicationUpdatedAt
		break
	case FieldPublicationDeletedAt:
		column = data_object.FieldPublicationDeletedAt
		break
	default:
		err = fmt.Errorf("unknown publication field: %s", field)
	}

	return column, skip, err
}

func (pi publicationInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy != nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldPublicationPublishedAt, DirDesc),
			fmt.Sprintf("%s %s", data_object.FieldPublicationCreatedAt, DirDesc),
		}
	}

	return args
}
