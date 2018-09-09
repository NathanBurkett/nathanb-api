package data_object_test

import (
	"testing"
	"time"

	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/satori/go.uuid"
)

func TestCategory_Table(t *testing.T) {
	type fields struct {
		ID           uuid.UUID
		Title        string
		Slug         string
		Publications []*data_object.Publication
		CreatedAt    time.Time
		UpdatedAt    time.Time
		DeletedAt    time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Passes",
			fields: fields{},
			want: data_object.TableCategory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := data_object.Category{
				ID:           tt.fields.ID,
				Title:        tt.fields.Title,
				Slug:         tt.fields.Slug,
				Publications: tt.fields.Publications,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				DeletedAt:    tt.fields.DeletedAt,
			}
			if got := c.Table(); got != tt.want {
				t.Errorf("Category.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
