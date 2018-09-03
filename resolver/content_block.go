package resolver

import (
	"github.com/nathanburkett/graphql-go"
	"context"
	"github.com/nathanburkett/nathanb-api/data_object"
)

type contentBlock struct {
	entity data_object.ContentBlock
}

func NewContentBlock(cb data_object.ContentBlock) *contentBlock {
	return &contentBlock{
		entity: cb,
	}
}

func (cb *contentBlock) ID(ctx context.Context) graphql.ID {
	return graphql.ID(cb.entity.ID.String())
}

func (cb *contentBlock) Type(ctx context.Context) string {
	return cb.entity.Type
}

func (cb *contentBlock) Content(ctx context.Context) string {
	return cb.entity.Content
}

func (cb *contentBlock) Publication(ctx context.Context) *Publication {
	return &Publication{}
}

func (cb *contentBlock) CreatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: cb.entity.CreatedAt}
}

func (cb *contentBlock) UpdatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: cb.entity.UpdatedAt}
}

func (cb *contentBlock) DeletedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: cb.entity.DeletedAt}
}
