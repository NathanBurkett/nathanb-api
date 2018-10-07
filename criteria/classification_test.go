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

func TestClassification(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.Classification{},
				args: criteria.SingleClassificationArgs{
					ID: uuidPointer(),
					Title: stringPointer(),
					Slug: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldClassificationId,
					},
					{
						Name: criteria.FieldClassificationTitle,
					},
					{
						Name: criteria.FieldClassificationSlug,
					},
					{
						Name: criteria.FieldClassificationCreatedAt,
					},
					{
						Name: criteria.FieldClassificationUpdatedAt,
					},
					{
						Name: criteria.FieldClassificationDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s = ?",
					sliceToColumns(data_object.Classification{}.Fields()),
					data_object.TableClassification,
					data_object.FieldClassificationId,
					data_object.FieldClassificationTitle,
					data_object.FieldClassificationSlug,
				)
			}(),
			argsCount: 3,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.Classification{},
				args: criteria.SingleClassificationArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("classification table unknown classification field: foo"),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SingleClassificationArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.Classification{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldClassificationId,
					},
					{
						Name: criteria.FieldClassificationTitle,
					},
					{
						Name: criteria.FieldClassificationSlug,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s ORDER BY %s %s LIMIT %d",
					data_object.FieldClassificationId,
					data_object.FieldClassificationTitle,
					data_object.FieldClassificationSlug,
					data_object.TableClassification,
					data_object.FieldClassificationSlug,
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
				model: data_object.Classification{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown classification argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
