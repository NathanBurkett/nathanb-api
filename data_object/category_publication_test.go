package data_object_test

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/satori/go.uuid"
)

func TestCategoryPublication_Table(t *testing.T) {
	type fields struct {
		CategoryId    uuid.UUID
		PublicationId uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Passes",
			fields: fields{},
			want: data_object.TableCategoryPublication,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := data_object.CategoryPublication{
				CategoryId:    tt.fields.CategoryId,
				PublicationId: tt.fields.PublicationId,
			}
			if got := cp.Table(); got != tt.want {
				t.Errorf("CategoryPublication.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCategoryPublication_Fields(t *testing.T) {
	c := data_object.CategoryPublication{}
	assert.Equal(t, []string{
		data_object.FieldCategoryPublicationId,
		data_object.FieldCategoryPublicationCategoryId,
		data_object.FieldCategoryPublicationPublicationId,
	}, c.Fields())
}
