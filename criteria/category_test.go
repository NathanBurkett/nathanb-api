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

func TestCategory(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.Category{},
				args: criteria.SingleCategoryArgs{
					ID: uuidPointer(),
					Title: stringPointer(),
					Slug: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldCategoryId,
					},
					{
						Name: criteria.FieldCategoryTitle,
					},
					{
						Name: criteria.FieldCategorySlug,
					},
					{
						Name: criteria.FieldCategoryCreatedAt,
					},
					{
						Name: criteria.FieldCategoryUpdatedAt,
					},
					{
						Name: criteria.FieldCategoryDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s = ?",
					sliceToColumns(data_object.Category{}.Fields()),
					data_object.TableCategory,
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
				)
			}(),
			argsCount: 3,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.Category{},
				args: criteria.SingleCategoryArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       fmt.Errorf("%s table unknown field: foo", data_object.TableCategory),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SingleCategoryArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldCategoryId,
					},
					{
						Name: criteria.FieldCategoryTitle,
					},
					{
						Name: criteria.FieldCategorySlug,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategorySlug,
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
				model: data_object.Category{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown category argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
