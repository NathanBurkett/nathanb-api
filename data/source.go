package data

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Database interface {
	Get(interface{}, string, ...interface{}) error
	Select(interface{}, string, ...interface{}) error
}

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
