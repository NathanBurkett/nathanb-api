package criteria

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	"fmt"
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

type mediaInterpretation struct{}

func (mi mediaInterpretation) handleArgs(c AbstractCriteria, args interface{}) {

}

func (mi mediaInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err error
		skip bool
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
