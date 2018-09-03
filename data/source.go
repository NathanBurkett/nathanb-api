package data

import (
	"log"
	"github.com/jmoiron/sqlx"
)

type Source struct {
	dsn string
	db  *sqlx.DB
}

func NewSource(dsn string) *Source {
	return &Source{
		dsn: dsn,
	}
}

func(s *Source) Connect() *Source {
	if s.db != nil {
		return s
	}

	conn, err := sqlx.Open("mysql", s.dsn)
	if err != nil {
		log.Fatal(err)
	}

	s.db = conn

	return s
}

func (s *Source) DB() *sqlx.DB {
	return s.db
}
