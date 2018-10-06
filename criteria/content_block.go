package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

const FieldContentBlockID = FieldId
const FieldContentBlockType = "type"
const FieldContentBlockContent = "content"
const FieldContentBlockOrder = "order"
const FieldContentBlockPublication = "publication"
const FieldContentBlockCreatedAt = FieldCreatedAt
const FieldContentBlockUpdatedAt = FieldUpdatedAt
const FieldContentBlockDeletedAt = FieldDeletedAt

type FirstContentBlockArgs struct {
	ID   *uuid.UUID
	Type *string
}

type contentBlockInterpretation struct{}

func (cbi contentBlockInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	if c.Error() != nil {
		return
	}

	switch T := args.(type) {
	case FirstContentBlockArgs:
		cbi.interpretFirstContentBlockArgs(c, args)
		break
	case PaginationArgs:
		T = cbi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown content block argument type: %s", T))
	}

}

func (cbi contentBlockInterpretation) interpretFirstContentBlockArgs(c AbstractCriteria, args interface{}) {
	firstArgs := args.(FirstContentBlockArgs)

	if firstArgs.ID != nil {
		c.Where(query.Eq{data_object.FieldContentBlockID: firstArgs.ID})
	}

	if firstArgs.Type != nil {
		c.Where(query.Eq{data_object.FieldContentBlockType: firstArgs.Type})
	}
}

func (cbi contentBlockInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err    error
		skip   bool
	)

	switch field {
	case FieldContentBlockID:
		column = data_object.FieldContentBlockID
		break
	case FieldContentBlockType:
		column = data_object.FieldContentBlockType
		break
	case FieldContentBlockContent:
		column = data_object.FieldContentBlockContent
		break
	case FieldContentBlockOrder:
		column = data_object.FieldContentBlockOrder
		break
	case FieldContentBlockCreatedAt:
		column = data_object.FieldContentBlockCreatedAt
		break
	case FieldContentBlockUpdatedAt:
		column = data_object.FieldContentBlockUpdatedAt
		break
	case FieldContentBlockDeletedAt:
		column = data_object.FieldContentBlockDeletedAt
		break
	default:
		err = fmt.Errorf("unknown content block field: %s", field)
	}

	return column, skip, err
}

func (cbi contentBlockInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	return args
}
