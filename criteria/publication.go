package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldPublicationId = FieldId
const FieldPublicationTitle = "title"
const FieldPublicationSlug = "slug"
const FieldPublicationPublishedAt = "publishedAt"
const FieldPublicationCreatedAt = FieldCreatedAt
const FieldPublicationUpdatedAt = FieldUpdatedAt
const FieldPublicationDeletedAt = FieldDeletedAt

type SinglePublicationArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type publicationInterpretation struct{}

func (publicationInterpretation) fields() map[string]string {
	return map[string]string{
		FieldPublicationId:          data_object.FieldPublicationId,
		FieldPublicationTitle:       data_object.FieldPublicationTitle,
		FieldPublicationSlug:        data_object.FieldPublicationSlug,
		FieldPublicationPublishedAt: data_object.FieldPublicationPublishedAt,
		FieldPublicationCreatedAt:   data_object.FieldPublicationCreatedAt,
		FieldPublicationUpdatedAt:   data_object.FieldPublicationUpdatedAt,
		FieldPublicationDeletedAt:   data_object.FieldPublicationDeletedAt,
	}
}

func (pi publicationInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SinglePublicationArgs:
		pi.interpretFirstPublicationArgs(c, T)
		break
	case PaginationArgs:
		T = pi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown publication argument type: %s", reflect.TypeOf(T)))
	}
}

func (pi publicationInterpretation) interpretFirstPublicationArgs(criteria AbstractCriteria, args SinglePublicationArgs) {
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
		err    error
		skip   bool
	)

	column = pi.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown field: %s", field)
	}

	return column, skip, err
}

func (pi publicationInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldPublicationPublishedAt, DirDesc),
			fmt.Sprintf("%s %s", data_object.FieldPublicationCreatedAt, DirDesc),
		}
	}

	return args
}
