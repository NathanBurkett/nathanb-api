package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nathanburkett/graphql-go"
	"github.com/nathanburkett/nathanb-api/app"
	"github.com/nathanburkett/nathanb-api/data"
	"github.com/nathanburkett/nathanb-api/env"
	"github.com/nathanburkett/nathanb-api/resolver"
	"github.com/nathanburkett/nathanb-api/schema"
	"github.com/nathanburkett/nathanb-api/schema_standard"
	"log"
	"os"
)

func main() {
	instance := app.NewInstance()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	instance.SetRootDir(pwd)
	env.ReadEnv(instance)

	dataSource := data.NewSource(env.Must("DB_DSN")).Connect()

	instance.SetDataSource(dataSource)
	instance.SetSchema(parseSchema(schema_standard.Definition{}, dataSource.DB()))
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
