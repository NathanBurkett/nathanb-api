package resolver

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	"context"
	"github.com/nathanburkett/graphql-go"
)

type category struct {
	entity data_object.Category
}

func NewCategoryResolver(cat data_object.Category) *category {
	return &category{
		entity: cat,
	}
}

func (c *category) ID(ctx context.Context) graphql.ID {
	return graphql.ID(c.entity.ID.String())
}

func (c *category) Title(ctx context.Context) string {
	return c.entity.Title
}

func (c *category) Slug(ctx context.Context) string {
	return c.entity.Slug
}

func (c *category) Publications(ctx context.Context) []*Publication {
	var pubs []*Publication

	for i := 0; i < len(c.entity.Publications); i++ {
		pubs = append(pubs, NewPublication(c.entity.Publications[i]))
	}

	return pubs
}

func (c *category) CreatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: c.entity.CreatedAt}
}
func (c *category) UpdatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: c.entity.UpdatedAt}
}
func (c *category) DeletedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: c.entity.DeletedAt}
}
