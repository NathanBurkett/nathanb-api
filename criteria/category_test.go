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
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause", criteria.QualifierGT),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategoryCreatedAt: {
							{
								Qualifier: func () *string { v := criteria.QualifierGT; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "2018-01-01"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s > ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategoryCreatedAt,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause", criteria.QualifierGTE),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategoryCreatedAt: {
							{
								Qualifier: func () *string { v := criteria.QualifierGTE; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "2018-01-01"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s >= ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategoryCreatedAt,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause", criteria.QualifierLT),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategoryCreatedAt: {
							{
								Qualifier: func () *string { v := criteria.QualifierLT; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "2018-01-01"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s < ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategoryCreatedAt,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause", criteria.QualifierLTE),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategoryCreatedAt: {
							{
								Qualifier: func () *string { v := criteria.QualifierLTE; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "2018-01-01"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s <= ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategoryCreatedAt,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause", criteria.QualifierEq),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategorySlug: {
							{
								Qualifier: func () *string { v := criteria.QualifierEq; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "foo"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s = ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategorySlug,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause", criteria.QualifierNotEq),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategorySlug: {
							{
								Qualifier: func () *string { v := criteria.QualifierNotEq; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "foo"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s <> ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategorySlug,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s where clause w/o actual qualifier", criteria.QualifierEq),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						criteria.FieldCategorySlug: {
							{
								Qualifier: func () *string { v := ""; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "foo"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s WHERE %s = ? ORDER BY %s %s LIMIT %d",
					data_object.FieldCategoryId,
					data_object.FieldCategoryTitle,
					data_object.FieldCategorySlug,
					data_object.TableCategory,
					data_object.FieldCategorySlug,
					data_object.FieldCategorySlug,
					criteria.DirDesc,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 1,
			err:       nil,
		},
		{
			name: fmt.Sprintf("Uses PaginationArgs w/ %s qualifier w/o valid column returns error", criteria.QualifierEq),
			args: testArgs{
				model: data_object.Category{},
				args: criteria.PaginationArgs{
					Where: &map[string][]criteria.WhereClause{
						"foo": {
							{
								Qualifier: func () *string { v := ""; return &v }(),
								Value: func () *interface{} { var v interface{}; v = "foo"; return &v }(),
							},
						},
					},
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
				},
			},
			wantSql: "",
			argsCount: 0,
			err:       errors.New("unknown field: foo"),
		},
	}

	RunCriteriaTests(t, tests)
}
