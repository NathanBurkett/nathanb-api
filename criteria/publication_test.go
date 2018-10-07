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

func TestPublication(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.Publication{},
				args: criteria.SinglePublicationArgs{
					ID: uuidPointer(),
					Title: stringPointer(),
					Slug: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldPublicationId,
					},
					{
						Name: criteria.FieldPublicationTitle,
					},
					{
						Name: criteria.FieldPublicationSlug,
					},
					{
						Name: criteria.FieldPublicationPublishedAt,
					},
					{
						Name: criteria.FieldPublicationCreatedAt,
					},
					{
						Name: criteria.FieldPublicationUpdatedAt,
					},
					{
						Name: criteria.FieldPublicationDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s = ?",
					sliceToColumns(data_object.Publication{}.Fields()),
					data_object.TablePublication,
					data_object.FieldPublicationId,
					data_object.FieldPublicationTitle,
					data_object.FieldPublicationSlug,
				)
			}(),
			argsCount: 3,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.Publication{},
				args: criteria.SinglePublicationArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       fmt.Errorf("%s table unknown field: foo", data_object.TablePublication),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SinglePublicationArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.Publication{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldPublicationId,
					},
					{
						Name: criteria.FieldPublicationTitle,
					},
					{
						Name: criteria.FieldPublicationSlug,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s ORDER BY %s %s, %s %s LIMIT %d",
					data_object.FieldPublicationId,
					data_object.FieldPublicationTitle,
					data_object.FieldPublicationSlug,
					data_object.TablePublication,
					data_object.FieldPublicationPublishedAt,
					criteria.DirDesc,
					data_object.FieldPublicationCreatedAt,
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
				model: data_object.Publication{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown publication argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
