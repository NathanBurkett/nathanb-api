package resolver

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/nathanburkett/graphql-go/selected"
	"github.com/nathanburkett/nathanb-api/repository"
	"github.com/nathanburkett/nathanb-api/criteria"
	"github.com/nathanburkett/nathanb-api/data_object"
	"database/sql"
)

type Query struct {
	DB         *sqlx.DB
}

func NewQuery(db *sqlx.DB) *Query {
	return &Query{
		DB:         db,
	}
}

func (q *Query) Category(ctx context.Context, args criteria.FirstCategoryArgs, fields []selected.SelectedField) (*category, error) {
	var cat data_object.Category

	cri := criteria.New(cat, args, criteria.ExtractFieldsFromSelectedFields(fields))
	err := repository.New(q.DB).First(&cat, cri)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return NewCategoryResolver(cat), nil
}

func (q *Query) Categories(ctx context.Context, args criteria.PaginationArgs, fields []selected.SelectedField) (*[]*category, error) {
	var cats []data_object.Category

	cri := criteria.New(data_object.Category{}, args, criteria.ExtractFieldsFromSelectedFields(fields))
	err := repository.New(q.DB).All(&cats, cri)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	var vals []*category

	for i := 0; i < len(cats); i++ {
		vals = append(vals, NewCategoryResolver(cats[i]))
	}

	return &vals, nil
}

func (q *Query) Classification(ctx context.Context) {}

func (q *Query) Classifications(ctx context.Context) {}

func (q *Query) ContentBlock(ctx context.Context) {}

func (q *Query) ContentBlocks(ctx context.Context) {}

func (q *Query) Media(ctx context.Context) {}

func (q *Query) Medias(ctx context.Context) {}

func (q *Query) Profile(ctx context.Context) {}

func (q *Query) Profiles(ctx context.Context) {}

func (q *Query) Publication(ctx context.Context) {}

func (q *Query) Publications(ctx context.Context) {}

func (q *Query) User(ctx context.Context) {}

func (q *Query) Users(ctx context.Context) {}
