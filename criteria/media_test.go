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

func TestMedia(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.Media{},
				args: criteria.SingleMediaArgs{
					ID: uuidPointer(),
					Path: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldMediaId,
					},
					{
						Name: criteria.FieldMediaType,
					},
					{
						Name: criteria.FieldMediaSubtype,
					},
					{
						Name: criteria.FieldMediaTitle,
					},
					{
						Name: criteria.FieldMediaPath,
					},
					{
						Name: criteria.FieldMediaAlt,
					},
					{
						Name: criteria.FieldMediaAspectRatio,
					},
					{
						Name: criteria.FieldMediaCreatedAt,
					},
					{
						Name: criteria.FieldMediaUpdatedAt,
					},
					{
						Name: criteria.FieldMediaDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ?",
					sliceToColumns(data_object.Media{}.Fields()),
					data_object.TableMedia,
					data_object.FieldMediaId,
					data_object.FieldMediaPath,
				)
			}(),
			argsCount: 2,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.Media{},
				args: criteria.SingleMediaArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       fmt.Errorf("%s table unknown field: foo", data_object.TableMedia),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SingleMediaArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.Media{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldMediaId,
					},
					{
						Name: criteria.FieldMediaTitle,
					},
					{
						Name: criteria.FieldMediaPath,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s LIMIT %d",
					data_object.FieldMediaId,
					data_object.FieldMediaTitle,
					data_object.FieldMediaPath,
					data_object.TableMedia,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 0,
			err:       nil,
		},
		{
			name: "Want err - has unknown args",
			args: testArgs{
				model: data_object.Media{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown media argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
