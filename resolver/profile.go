package resolver

import (
	"github.com/nathanburkett/graphql-go"
	"github.com/nathanburkett/nathanb-api/data_object"
	"context"
	"bytes"
)

type profile struct {
	entity data_object.Profile
}

func NewProfile() *profile {
	return &profile{}
}

func (p *profile) ID(ctx context.Context) graphql.ID {
	return graphql.ID(p.entity.ID.String())
}

func (p *profile) FirstName(ctx context.Context) string {
	return p.entity.FirstName
}

func (p *profile) LastName(ctx context.Context) string {
	return p.entity.LastName
}

func (p *profile) Name(ctx context.Context) string {
	buf := bytes.Buffer{}

	if p.entity.FirstName != "" {
		buf.WriteString(p.entity.FirstName)
	}

	if p.entity.LastName != "" {
		if buf.Len() > 0 {
			buf.WriteString(" ")
		}

		buf.WriteString(p.entity.LastName)
	}
	return buf.String()
}

func (p *profile) TwitterHandle(ctx context.Context) string {
	return p.entity.TwitterHandle
}

func (p *profile) GithubHandle(ctx context.Context) string {
	return p.entity.GithubHandle
}

//func (p *profile) User(ctx context.Context) *User {
//	return NewUser()
//}

func (p *profile) CreatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.CreatedAt}
}

func (p *profile) UpdatedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.UpdatedAt}
}

func (p *profile) DeletedAt(ctx context.Context) graphql.Time {
	return graphql.Time{Time: p.entity.DeletedAt}
}
