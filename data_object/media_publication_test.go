package data_object_test

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/satori/go.uuid"
)

func TestMediaPublication_Table(t *testing.T) {
	type fields struct {
		MediaId       uuid.UUID
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
			want: data_object.TableMediaPublication,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := data_object.MediaPublication{
				MediaId:       tt.fields.MediaId,
				PublicationId: tt.fields.PublicationId,
			}
			if got := mp.Table(); got != tt.want {
				t.Errorf("MediaPublication.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMediaPublication_Fields(t *testing.T) {
	c := data_object.MediaPublication{}
	assert.Equal(t, []string{
		data_object.FieldMediaPublicationId,
		data_object.FieldMediaPublicationMediaId,
		data_object.FieldMediaPublicationPublicationId,
	}, c.Fields())
}
