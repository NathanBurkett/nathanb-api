package resolver

import (
	"github.com/nathanburkett/graphql-go"
	"github.com/nathanburkett/nathanb-api/data_object"
	"context"
)

type Publication struct {
	entity *data_object.Publication
}

func NewPublication(pub *data_object.Publication) *Publication {
	return &Publication{
		entity: pub,
	}
}

func (p *Publication) ID(ctx context.Context) graphql.ID {
	return graphql.ID(p.entity.ID.String())
}

func (p *Publication) Title(ctx context.Context) string {
	return p.entity.Title
}

func (p *Publication) Slug(ctx context.Context) string {
	return p.entity.Slug
}

//func (p *Publication) Classification(ctx context.Context) *Classification {
//	return p.entity.Classification
//}

//func (p *Publication) Categories(ctx context.Context) []*data_object.category {
//	return p.entity.Categories
//}
//
//func (p *Publication) ContentBlocks(ctx context.Context) []*data_object.contentBlock {
//	return p.entity.ContentBlocks
//}
//
//func (p *Publication) media(ctx context.Context) []*data_object.media {
//	return p.entity.media
//}
//
//func (p *Publication) Users(ctx context.Context) []*data_object.User {
//	return p.entity.Users
//}

func (p *Publication) PublishedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.PublishedAt}
}

func (p *Publication) CreatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.CreatedAt}
}

func (p *Publication) UpdatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.UpdatedAt}
}
func (p *Publication) DeletedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.DeletedAt}
}
