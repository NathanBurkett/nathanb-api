package criteria_test

import (
	"errors"
	"fmt"
	"github.com/nathanburkett/graphql-go/selected"
	"github.com/nathanburkett/nathanb-api/criteria"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/nathanburkett/nathanb-api/mock"
	"testing"
)

func TestContentBlock(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.ContentBlock{},
				args: criteria.SingleContentBlockArgs{
					ID: uuidPointer(),
					Type: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldContentBlockId,
					},
					{
						Name: criteria.FieldContentBlockType,
					},
					{
						Name: criteria.FieldContentBlockContent,
					},
					{
						Name: criteria.FieldContentBlockOrder,
					},
					{
						Name: criteria.FieldContentBlockCreatedAt,
					},
					{
						Name: criteria.FieldContentBlockUpdatedAt,
					},
					{
						Name: criteria.FieldContentBlockDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ?",
					sliceToColumns(data_object.ContentBlock{}.Fields()),
					data_object.TableContentBlock,
					data_object.FieldContentBlockId,
					data_object.FieldContentBlockType,
				)
			}(),
			argsCount: 2,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.ContentBlock{},
				args: criteria.SingleContentBlockArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("content_block table unknown content block field: foo"),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SingleContentBlockArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.ContentBlock{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldContentBlockId,
					},
					{
						Name: criteria.FieldContentBlockType,
					},
					{
						Name: criteria.FieldContentBlockContent,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s ORDER BY `%s` %s LIMIT %d",
					data_object.FieldContentBlockId,
					criteria.FieldContentBlockType,
					criteria.FieldContentBlockContent,
					data_object.TableContentBlock,
					criteria.FieldContentBlockOrder,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 0,
			err:       nil,
		},
		{
			name: "Want err - has unknown args",
			args: testArgs{
				model: data_object.ContentBlock{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown content block argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
