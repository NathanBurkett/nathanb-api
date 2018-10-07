package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldMediaId = FieldId
const FieldMediaType = "type"
const FieldMediaSubtype = "subtype"
const FieldMediaTitle = "title"
const FieldMediaPath = "path"
const FieldMediaAlt = "alt"
const FieldMediaAspectRatio = "aspectRatio"
const FieldMediaCreatedAt = FieldCreatedAt
const FieldMediaUpdatedAt = FieldUpdatedAt
const FieldMediaDeletedAt = FieldDeletedAt
const FieldMediaPublications = "publications"

type SingleMediaArgs struct {
	ID   *uuid.UUID
	Path *string
}

type mediaInterpretation struct{}

func (mediaInterpretation) fields() map[string]string {
	return map[string]string{
		FieldMediaId:          data_object.FieldMediaId,
		FieldMediaType:        data_object.FieldMediaType,
		FieldMediaSubtype:     data_object.FieldMediaSubtype,
		FieldMediaTitle:       data_object.FieldMediaTitle,
		FieldMediaPath:        data_object.FieldMediaPath,
		FieldMediaAlt:         data_object.FieldMediaAlt,
		FieldMediaAspectRatio: data_object.FieldMediaAspectRatio,
		FieldMediaCreatedAt:   data_object.FieldMediaCreatedAt,
		FieldMediaUpdatedAt:   data_object.FieldMediaUpdatedAt,
		FieldMediaDeletedAt:   data_object.FieldMediaDeletedAt,
	}
}

func (mi mediaInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SingleMediaArgs:
		mi.interpretFirstCategoryArgs(c, T)
		break
	case PaginationArgs:
		T = mi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown media argument type: %s", reflect.TypeOf(T)))
	}
}

func (mi mediaInterpretation) interpretFirstCategoryArgs(criteria AbstractCriteria, args SingleMediaArgs) {
	if args.ID != nil {
		criteria.Where(query.Eq{data_object.FieldMediaId: args.ID})
	}

	if args.Path != nil {
		criteria.Where(query.Eq{data_object.FieldMediaPath: args.Path})
	}
}

func (mi mediaInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err    error
		skip   bool
	)

	column = mi.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown field: %s", field)
	}

	return column, skip, err
}

func (mi mediaInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{}
	}

	return args
}
