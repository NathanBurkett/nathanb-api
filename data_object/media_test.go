package data_object_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/nathanburkett/nathanb-api/data_object"
)

func TestMedia_Table(t *testing.T) {
	type fields struct {
		ID           uuid.UUID
		Type         string
		Subtype      string
		Title        string
		Path         string
		Alt          string
		AspectRatio  string
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
			want: data_object.TableMedia,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := data_object.Media{
				ID:           tt.fields.ID,
				Type:         tt.fields.Type,
				Subtype:      tt.fields.Subtype,
				Title:        tt.fields.Title,
				Path:         tt.fields.Path,
				Alt:          tt.fields.Alt,
				AspectRatio:  tt.fields.AspectRatio,
				Publications: tt.fields.Publications,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				DeletedAt:    tt.fields.DeletedAt,
			}
			if got := m.Table(); got != tt.want {
				t.Errorf("Media.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedia_Fields(t *testing.T) {
	c := data_object.Media{}
	assert.Equal(t, []string{
		data_object.FieldMediaId,
		data_object.FieldMediaType,
		data_object.FieldMediaSubtype,
		data_object.FieldMediaTitle,
		data_object.FieldMediaPath,
		data_object.FieldMediaAlt,
		data_object.FieldMediaAspectRatio,
		data_object.FieldMediaCreatedAt,
		data_object.FieldMediaUpdatedAt,
		data_object.FieldMediaDeletedAt,
	}, c.Fields())
}
