package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/nathanburkett/nathanb-api/app"
	"github.com/nathanburkett/nathanb-api/env"
	"fmt"
	"github.com/nathanburkett/nathanb-api/schema"
	"github.com/nathanburkett/nathanb-api/data"
	"github.com/nathanburkett/nathanb-api/schema_standard"
	"github.com/jmoiron/sqlx"
)

func main() {
	dataSource := data.NewSource(env.Must("DB_DSN")).Connect()
	instance := app.NewInstance()

	instance.SetDataSource(dataSource)

	standard := schema_standard.Definition{}
	instance.SetSchema(parseSchema(standard, dataSource.DB()))
}

func getHostAndPort() string {
	return fmt.Sprintf("%s:%s", env.Must("APP_HOST"), env.Must("APP_PORT"))
}

func parseSchema(def schema.Definition, db *sqlx.DB) *graphql.Schema {
	s, err := schema.NewParser().Parse(def.Define(), resolver.NewQuery(db))
	if err != nil {
		log.Fatal(err)
	}

	return s
}
