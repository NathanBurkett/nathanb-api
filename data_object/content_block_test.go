package data_object_test

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/nathanburkett/nathanb-api/data_object"
)

func TestContentBlock_Table(t *testing.T) {
	type fields struct {
		ID            uuid.UUID
		Type          string
		Content       string
		Order         uint8
		Publication   data_object.Publication
		PublicationId uuid.UUID
		CreatedAt     time.Time
		UpdatedAt     time.Time
		DeletedAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Passes",
			fields: fields{},
			want: data_object.TableContentBlock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := data_object.ContentBlock{
				ID:            tt.fields.ID,
				Type:          tt.fields.Type,
				Content:       tt.fields.Content,
				Order:         tt.fields.Order,
				Publication:   tt.fields.Publication,
				PublicationId: tt.fields.PublicationId,
				CreatedAt:     tt.fields.CreatedAt,
				UpdatedAt:     tt.fields.UpdatedAt,
				DeletedAt:     tt.fields.DeletedAt,
			}
			if got := cb.Table(); got != tt.want {
				t.Errorf("ContentBlock.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
