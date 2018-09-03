package criteria

import (
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/nathanburkett/nathanb-api/data_object"
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

func (cl classificationInterpretation) handleArgs(c *Criteria, args interface{}) {
	if c.err != nil {
		return
	}

	switch T := args.(type) {
	case FirstClassificationArgs:
		break
	case PaginationArgs:
		break
	default:
		c.err = fmt.Errorf("unknown classification argument type: %s", T)
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
