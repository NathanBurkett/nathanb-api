package data_object_test

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/nathanburkett/nathanb-api/data_object"
)

func TestClassification_Table(t *testing.T) {
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
			want: data_object.TableClassification,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := data_object.Classification{
				ID:           tt.fields.ID,
				Title:        tt.fields.Title,
				Slug:         tt.fields.Slug,
				Publications: tt.fields.Publications,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				DeletedAt:    tt.fields.DeletedAt,
			}
			if got := c.Table(); got != tt.want {
				t.Errorf("Classification.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
