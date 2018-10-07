package data_object_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/nathanburkett/nathanb-api/data_object"
)

func TestUser_Table(t *testing.T) {
	type fields struct {
		ID             uuid.UUID
		Email          string
		PasswordDigest string
		Profile        data_object.Profile
		Publications   []data_object.Publication
		CreatedAt      time.Time
		UpdatedAt      time.Time
		DeletedAt      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Passes",
			fields: fields{},
			want: data_object.TableUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := data_object.User{
				ID:             tt.fields.ID,
				Email:          tt.fields.Email,
				PasswordDigest: tt.fields.PasswordDigest,
				Profile:        tt.fields.Profile,
				Publications:   tt.fields.Publications,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
				DeletedAt:      tt.fields.DeletedAt,
			}
			if got := u.Table(); got != tt.want {
				t.Errorf("User.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Fields(t *testing.T) {
	c := data_object.User{}
	assert.Equal(t, []string{
		data_object.FieldUserId,
		data_object.FieldUserEmail,
		data_object.FieldUserPasswordDigest,
		data_object.FieldUserCreatedAt,
		data_object.FieldUserUpdatedAt,
		data_object.FieldUserDeletedAt,
	}, c.Fields())
}
