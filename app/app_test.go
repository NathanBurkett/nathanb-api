package app_test

import (
	"github.com/nathanburkett/nathanb-api/app"
	"testing"

	"github.com/nathanburkett/graphql-go"
	"github.com/nathanburkett/nathanb-api/data"
	"github.com/stretchr/testify/assert"
)

func TestNewInstance(t *testing.T) {
	tests := []struct {
		name string
		want *app.Instance
	}{
		{
			name: "Valid Instance",
			want: &app.Instance{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := app.NewInstance()

			assert.NotNil(t, got, "app.NewInstance() creates new app.Instance")
			assert.Nil(t, got.DataSource, "app.Instance.DataSource is nil")
			assert.Nil(t, got.Schema, "app.Instance.Schema is nil")
		})
	}
}

func TestInstance_SetDataSource(t *testing.T) {
	type fields struct {
		DataSource *data.Source
		Schema     *graphql.Schema
	}
	type args struct {
		ds *data.Source
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *app.Instance
	}{
		{
			name: "Provides *data.Source",
			fields: fields{},
			args: args{
				ds: &data.Source{},
			},
			want: &app.Instance{
				DataSource: &data.Source{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &app.Instance{
				DataSource: tt.fields.DataSource,
				Schema:     tt.fields.Schema,
			}

			i.SetDataSource(tt.args.ds)

			assert.NotNil(t, i.DataSource, "Instance.SetDataSource() correctly sets Instance.DataSource")
		})
	}
}

func TestInstance_SetSchema(t *testing.T) {
	type fields struct {
		DataSource *data.Source
		Schema     *graphql.Schema
	}
	type args struct {
		schema *graphql.Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *app.Instance
	}{
		{
			name: "Provides *data.Source",
			fields: fields{},
			args: args{
				schema: &graphql.Schema{},
			},
			want: &app.Instance{
				Schema: &graphql.Schema{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &app.Instance{
				DataSource: tt.fields.DataSource,
				Schema:     tt.fields.Schema,
			}

			i.SetSchema(tt.args.schema)

			assert.NotNil(t, i.Schema, "Instance.SetSchema() correctly sets Instance.Schema")
		})
	}
}
