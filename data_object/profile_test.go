package data_object_test

import (
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/nathanburkett/nathanb-api/data_object"
)

func TestProfile_Table(t *testing.T) {
	type fields struct {
		ID            uuid.UUID
		FirstName     string
		LastName      string
		TwitterHandle string
		GithubHandle  string
		UserID        uuid.UUID
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
			want: data_object.TableProfile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := data_object.Profile{
				ID:            tt.fields.ID,
				FirstName:     tt.fields.FirstName,
				LastName:      tt.fields.LastName,
				TwitterHandle: tt.fields.TwitterHandle,
				GithubHandle:  tt.fields.GithubHandle,
				UserID:        tt.fields.UserID,
				CreatedAt:     tt.fields.CreatedAt,
				UpdatedAt:     tt.fields.UpdatedAt,
				DeletedAt:     tt.fields.DeletedAt,
			}
			if got := p.Table(); got != tt.want {
				t.Errorf("Profile.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}
