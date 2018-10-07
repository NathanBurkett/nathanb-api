package criteria_test

import (
	"fmt"
	"strings"
	"testing"

	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/graphql-go/selected"
	"github.com/nathanburkett/nathanb-api/criteria"
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func uuidPointer() *uuid.UUID {
	u := uuid.NewV4()
	return &u
}

func stringPointer() *string {
	s := ""
	return &s
}

func uint64Pointer() *uint64 {
	s := uint64(1)
	return &s
}

func sliceToColumns(fieldSlice []string) string {
	return strings.Join(fieldSlice, ", ")
}

type testArgs struct {
	model  data_object.Model
	args   interface{}
	fields []selected.SelectedField
}

type testCase struct {
	name      string
	args      testArgs
	wantSql   string
	argsCount int
	err       error
}

func RunCriteriaTests(t *testing.T, testCases []testCase) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cri := criteria.New(tt.args.model, tt.args.args, tt.args.fields)

			if tt.err != nil && cri.Error().Error() != tt.err.Error() {
				t.Errorf("expected criteria.Error() '%v' but got '%v'", tt.err, cri.Error())
				return
			}

			str, i, err := cri.ToSql()

			if err != nil {
				assert.Equal(t, tt.err, err)
			}

			assert.Equal(t, tt.wantSql, str)
			assert.Equal(t, tt.argsCount, len(i))
		})
	}
}

func TestCriteria_Offset(t *testing.T) {
	type fields struct {
		builder     query.SelectBuilder
		modelType   data_object.Model
		interpreter criteria.ModelInterpreter
		err         error
	}
	type args struct {
		args   interface{}
		fields []selected.SelectedField
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantSql string
	}{
		{
			name: "Successfully sets Offset",
			fields: fields{
				modelType: data_object.Category{},
			},
			args: args{
				args: criteria.PaginationArgs{
					Page: func() *uint64 {
						v := uint64(2)
						return &v
					}(),
				},
				fields: []selected.SelectedField{
					{
						Name: criteria.FieldCategoryId,
					},
				},
			},
			wantSql: fmt.Sprintf(
				"SELECT %s FROM %s ORDER BY %s %s LIMIT 10 OFFSET 10",
				data_object.FieldCategoryId,
				data_object.TableCategory,
				data_object.FieldCategorySlug,
				criteria.DirDesc,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, _, err := criteria.New(tt.fields.modelType, tt.args.args, tt.args.fields).ToSql()
			assert.Equal(t, tt.wantSql, sql)
			assert.Nil(t, err)
		})
	}
}

func TestCriteria_Interpreter(t *testing.T) {
	type fields struct {
		builder     query.SelectBuilder
		modelType   data_object.Model
		interpreter criteria.ModelInterpreter
		err         error
	}
	type args struct {
		args   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args args
	}{
		{
			name: "Interpreter() returns successfully",
			fields: fields{
				modelType: data_object.Category{},
			},
			args: args{
				args: criteria.PaginationArgs{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := criteria.New(tt.fields.modelType, tt.args, []selected.SelectedField{})
			assert.NotNil(t, c.Interpreter())
		})
	}
}
