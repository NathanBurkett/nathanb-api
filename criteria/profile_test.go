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

func TestProfile(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.Profile{},
				args: criteria.SingleProfileArgs{
					ID: uuidPointer(),
					UserId: uint64Pointer(),
					GithubHandle: stringPointer(),
					TwitterHandle: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldProfileId,
					},
					{
						Name: criteria.FieldProfileFirstName,
					},
					{
						Name: criteria.FieldProfileLastName,
					},
					{
						Name: criteria.FieldProfileTwitterHandle,
					},
					{
						Name: criteria.FieldProfileGithubHandle,
					},
					{
						Name: criteria.FieldProfileUserId,
					},
					{
						Name: criteria.FieldProfileCreatedAt,
					},
					{
						Name: criteria.FieldProfileUpdatedAt,
					},
					{
						Name: criteria.FieldProfileDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s = ? AND %s = ?",
					sliceToColumns(data_object.Profile{}.Fields()),
					data_object.TableProfile,
					data_object.FieldProfileId,
					data_object.FieldProfileUserId,
					data_object.FieldProfileGithubHandle,
					data_object.FieldProfileTwitterHandle,
				)
			}(),
			argsCount: 4,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.Profile{},
				args: criteria.SingleProfileArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       fmt.Errorf("%s table unknown field: foo", data_object.TableProfile),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SingleProfileArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.Profile{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldProfileId,
					},
					{
						Name: criteria.FieldProfileFirstName,
					},
					{
						Name: criteria.FieldProfileLastName,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s, %s FROM %s ORDER BY %s %s, %s %s LIMIT %d",
					data_object.FieldProfileId,
					data_object.FieldProfileFirstName,
					data_object.FieldProfileLastName,
					data_object.TableProfile,
					data_object.FieldProfileLastName,
					criteria.DirDesc,
					data_object.FieldProfileFirstName,
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
				model: data_object.Profile{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown profile argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
