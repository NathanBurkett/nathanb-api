package criteria

import (
	"fmt"
	"github.com/nathanburkett/nathanb-api/data_object"
)

const FieldPublicationId = FieldId
const FieldPublicationTitle = "title"
const FieldPublicationSlug = "slug"
const FieldPublicationPublishedAt = "publishedAt"
const FieldPublicationCreatedAt = FieldCreatedAt
const FieldPublicationUpdatedAt = FieldUpdatedAt
const FieldPublicationDeletedAt = FieldDeletedAt

type publicationInterpretation struct {}

func (pi publicationInterpretation) handleArgs(c *Criteria, args interface{}) {

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
