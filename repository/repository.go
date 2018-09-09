package repository

import (
	"database/sql"
	"github.com/nathanburkett/nathanb-api/criteria"
	"github.com/nathanburkett/nathanb-api/data"
)

type Repository struct {
	DB data.Database
}

func New(db data.Database) Repository {
	return Repository{
		DB: db,
	}
}

func (r Repository) First(dest interface{}, cri criteria.AbstractCriteria) error {
	stmt, args, err := cri.ToSql()
	if err != nil {
		return err
	}

	if err := r.DB.Get(dest, stmt, args...); err != nil {
		return err
	}

	return nil
}

func (r Repository) Find(dest interface{}, cri criteria.AbstractCriteria) error {
	stmt, args, err := cri.ToSql()
	if err != nil {
		return err
	}

	if err := r.DB.Get(dest, stmt, args...); err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r Repository) All(dest interface{}, cri criteria.AbstractCriteria) error {
	stmt, args, err := cri.ToSql()
	if err != nil {
		return err
	}

	if err := r.DB.Select(dest, stmt, args...); err != nil {
		return err
	}

	return nil
}
