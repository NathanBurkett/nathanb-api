package data_object_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/satori/go.uuid"
	"github.com/nathanburkett/nathanb-api/data_object"
)

func TestPublication_Table(t *testing.T) {
	type fields struct {
		ID               uuid.UUID
		Title            string
		Slug             string
		ClassificationID uuid.UUID
		Classification   *data_object.Classification
		ContentBlocks    []*data_object.ContentBlock
		Categories       []*data_object.Category
		Media            []*data_object.Media
		Users            []*data_object.User
		PublishedAt      time.Time
		CreatedAt        time.Time
		UpdatedAt        time.Time
		DeletedAt        time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Passes",
			fields: fields{},
			want: data_object.TablePublication,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := data_object.Publication{
				ID:               tt.fields.ID,
				Title:            tt.fields.Title,
				Slug:             tt.fields.Slug,
				ClassificationID: tt.fields.ClassificationID,
				Classification:   tt.fields.Classification,
				ContentBlocks:    tt.fields.ContentBlocks,
				Categories:       tt.fields.Categories,
				Media:            tt.fields.Media,
				Users:            tt.fields.Users,
				PublishedAt:      tt.fields.PublishedAt,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				DeletedAt:        tt.fields.DeletedAt,
			}
			if got := p.Table(); got != tt.want {
				t.Errorf("Publication.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublication_Fields(t *testing.T) {
	c := data_object.Publication{}
	assert.Equal(t, []string{
		data_object.FieldPublicationId,
		data_object.FieldPublicationTitle,
		data_object.FieldPublicationSlug,
		data_object.FieldPublicationPublishedAt,
		data_object.FieldPublicationCreatedAt,
		data_object.FieldPublicationUpdatedAt,
		data_object.FieldPublicationDeletedAt,
	}, c.Fields())
}
