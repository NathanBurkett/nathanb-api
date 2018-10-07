package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"reflect"
)

const FieldContentBlockId = FieldId
const FieldContentBlockType = "type"
const FieldContentBlockContent = "content"
const FieldContentBlockOrder = "order"
const FieldContentBlockPublication = "publication"
const FieldContentBlockCreatedAt = FieldCreatedAt
const FieldContentBlockUpdatedAt = FieldUpdatedAt
const FieldContentBlockDeletedAt = FieldDeletedAt

type SingleContentBlockArgs struct {
	ID   *uuid.UUID
	Type *string
}

type contentBlockInterpretation struct{}

func (cbi contentBlockInterpretation) fields() map[string]string {
	return map[string]string{
		FieldContentBlockId:        data_object.FieldContentBlockId,
		FieldContentBlockType:      data_object.FieldContentBlockType,
		FieldContentBlockContent:   data_object.FieldContentBlockContent,
		FieldContentBlockOrder:     data_object.FieldContentBlockOrder,
		FieldContentBlockCreatedAt: data_object.FieldContentBlockCreatedAt,
		FieldContentBlockUpdatedAt: data_object.FieldContentBlockUpdatedAt,
		FieldContentBlockDeletedAt: data_object.FieldContentBlockDeletedAt,
	}
}

func (cbi contentBlockInterpretation) handleArgs(c AbstractCriteria, args interface{}) {
	switch T := args.(type) {
	case SingleContentBlockArgs:
		cbi.interpretFirstContentBlockArgs(c, args)
		break
	case PaginationArgs:
		T = cbi.checkDefaultPaginationArgs(T)
		interpretPaginationArgs(c, T)
		break
	default:
		c.SetError(fmt.Errorf("unknown content block argument type: %s", reflect.TypeOf(T)))
	}

}

func (cbi contentBlockInterpretation) interpretFirstContentBlockArgs(c AbstractCriteria, args interface{}) {
	firstArgs := args.(SingleContentBlockArgs)

	if firstArgs.ID != nil {
		c.Where(query.Eq{data_object.FieldContentBlockId: firstArgs.ID})
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

	column = cbi.fields()[field]
	if column == "" {
		err = fmt.Errorf("unknown content block field: %s", field)
	}

	return column, skip, err
}

func (cbi contentBlockInterpretation) checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	args = checkDefaultPaginationArgs(args)

	if args.OrderBy == nil {
		args.OrderBy = &[]string{
			fmt.Sprintf("%s %s", fmt.Sprintf("`%s`", data_object.FieldContentBlockOrder), DirDesc),
		}
	}

	return args
}
