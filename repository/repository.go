package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nathanburkett/nathanb-api/criteria"
	"database/sql"
)

type Repository struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{
		DB: db,
	}
}

func (r Repository) First(dest interface{}, cri *criteria.Criteria) error {
	stmt, args, err := cri.ToSql()
	if err != nil {
		return err
	}

	if err := r.DB.Get(dest, stmt, args...); err != nil {
		return err
	}

	return nil
}

func (r Repository) Find(dest interface{}, cri *criteria.Criteria) error {
	stmt, args, err := cri.ToSql()
	if err != nil {
		return err
	}

	if err := r.DB.Get(dest, stmt, args...); err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r Repository) All(dest interface{}, cri *criteria.Criteria) error {
	stmt, args, err := cri.ToSql()
	if err != nil {
		return err
	}

	if err := r.DB.Select(dest, stmt, args...); err != nil {
		return err
	}

	return nil
}
