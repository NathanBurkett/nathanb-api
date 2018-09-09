package data_object_test

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	"testing"

	"github.com/satori/go.uuid"
)

func TestPublicationUser_Table(t *testing.T) {
	type fields struct {
		PublicationId uuid.UUID
		UserId        uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Passes",
			fields: fields{},
			want: data_object.TablePublicationUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pu := data_object.PublicationUser{
				PublicationId: tt.fields.PublicationId,
				UserId:        tt.fields.UserId,
			}
			if got := pu.Table(); got != tt.want {
				t.Errorf("PublicationUser.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
