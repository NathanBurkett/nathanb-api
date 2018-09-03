package resolver

import (
	"github.com/nathanburkett/graphql-go"
	"context"
	"github.com/nathanburkett/nathanb-api/data_object"
)

type media struct {
	entity data_object.Media
}

func NewMedia() *media {
	return &media{}
}

func (m *media) ID(ctx context.Context) graphql.ID {
	return graphql.ID(m.entity.ID.String())
}

func (m *media) Type(ctx context.Context) string {
	return m.entity.Type
}

func (m *media) Subtype(ctx context.Context) string {
	return m.entity.Subtype
}

func (m *media) Title(ctx context.Context) string {
	return m.entity.Title
}

func (m *media) Path(ctx context.Context) string {
	return m.entity.Path
}

func (m *media) Alt(ctx context.Context) string {
	return m.entity.Alt
}

func (m *media) AspectRatio(ctx context.Context) string {
	return m.entity.AspectRatio
}

func (m *media) Publications(ctx context.Context) []*Publication {
	var pubs []*Publication

	for i := 0; i < len(m.entity.Publications); i++ {
		pubs = append(pubs, NewPublication(m.entity.Publications[i]))
	}

	return pubs
}

func (m *media) CreatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: m.entity.CreatedAt}
}

func (m *media) UpdatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: m.entity.UpdatedAt}
}

func (m *media) DeletedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: m.entity.DeletedAt}
}
