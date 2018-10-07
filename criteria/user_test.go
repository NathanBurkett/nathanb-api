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

func TestUser(t *testing.T) {
	tests := []testCase{
		{
			name: "Happy path",
			args: testArgs{
				model: data_object.User{},
				args: criteria.SingleUserArgs{
					ID: uuidPointer(),
					Email: stringPointer(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldUserId,
					},
					{
						Name: criteria.FieldUserEmail,
					},
					{
						Name: criteria.FieldUserPasswordDigest,
					},
					{
						Name: criteria.FieldUserCreatedAt,
					},
					{
						Name: criteria.FieldUserUpdatedAt,
					},
					{
						Name: criteria.FieldUserDeletedAt,
					},
				},
			},
			wantSql: func() string {
				return fmt.Sprintf(
					"SELECT %s FROM %s WHERE %s = ? AND %s = ?",
					sliceToColumns(data_object.User{}.Fields()),
					data_object.TableUser,
					data_object.FieldUserId,
					data_object.FieldUserEmail,
				)
			}(),
			argsCount: 2,
			err:       nil,
		},
		{
			name:      "Want err - unknown column",
			args:      testArgs{
				model: data_object.User{},
				args: criteria.SingleUserArgs{},
				fields: []selected.SelectedField{
					{
						Name: "foo",
					},
				},
			},
			wantSql:   "",
			argsCount: 0,
			err:       fmt.Errorf("%s table unknown field: foo", data_object.TableUser),
		},
		{
			name:      "Want err - unknown data object type",
			args:      testArgs{
				model: mock.Model{},
				args: criteria.SingleUserArgs{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown data object type: mock.Model"),
		},
		{
			name: "Uses default PaginationArgs",
			args: testArgs{
				model: data_object.User{},
				args: criteria.PaginationArgs{},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldUserId,
					},
					{
						Name: criteria.FieldUserEmail,
					},
				},
			},
			wantSql:   func() string {
				return fmt.Sprintf(
					"SELECT %s, %s FROM %s LIMIT %d",
					data_object.FieldUserId,
					data_object.FieldUserEmail,
					data_object.TableUser,
					criteria.DefaultLimit,
				)
			}(),
			argsCount: 0,
			err:       nil,
		},
		{
			name: "Want err - has unknown args",
			args: testArgs{
				model: data_object.User{},
				args: struct{}{},
				fields: []selected.SelectedField{},
			},
			wantSql:   "",
			argsCount: 0,
			err:       errors.New("unknown user argument type: struct {}"),
		},
	}

	RunCriteriaTests(t, tests)
}
