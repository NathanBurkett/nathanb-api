package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
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

type FirstMediaArgs struct {
	ID   *uuid.UUID
	Path *string
}

type mediaInterpretation struct{}

func (mi mediaInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstMediaArgs:
		mi.interpretFirstCategoryArgs(c, T)
		break
	case PaginationArgs:
		T = mi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown category argument type: %s", T))
	}
}

func (mi mediaInterpretation) interpretFirstCategoryArgs(criteria AbstractCriteria, args FirstMediaArgs) {
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

	switch field {
	case FieldMediaId:
		column = data_object.FieldMediaId
		break
	case FieldMediaType:
		column = data_object.FieldMediaType
		break
	case FieldMediaSubtype:
		column = data_object.FieldMediaSubtype
		break
	case FieldMediaTitle:
		column = data_object.FieldMediaTitle
		break
	case FieldMediaPath:
		column = data_object.FieldMediaPath
		break
	case FieldMediaAlt:
		column = data_object.FieldMediaAlt
		break
	case FieldMediaAspectRatio:
		column = data_object.FieldMediaAspectRatio
		break
	case FieldMediaCreatedAt:
		column = data_object.FieldMediaCreatedAt
		break
	case FieldMediaUpdatedAt:
		column = data_object.FieldMediaUpdatedAt
		break
	case FieldMediaDeletedAt:
		column = data_object.FieldMediaDeletedAt
		break
	case FieldMediaPublications:
		skip = true
		break
	default:
		err = fmt.Errorf("unknown media field: %s", field)
	}

	return column, skip, err
}

func (mi mediaInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy != nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", data_object.FieldCategoryCreatedAt, DirDesc),
		}
	}
	return args
}
