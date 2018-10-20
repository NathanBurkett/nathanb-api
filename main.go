package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nathanburkett/env"
	"github.com/nathanburkett/graphql-go"
	"github.com/nathanburkett/nathanb-api/app"
	"github.com/nathanburkett/nathanb-api/data"
	"github.com/nathanburkett/nathanb-api/resolver"
	"github.com/nathanburkett/nathanb-api/schema"
	"github.com/nathanburkett/nathanb-api/schema_standard"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	instance := app.NewInstance().SetRootDir(pwd)
	readEnv(instance)

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

func readEnv(instance *app.Instance) {
	envPath := fmt.Sprintf("%s/.env", instance.RootDir())

	file, err := os.Open(envPath)
	defer file.Close()
	if err != nil {
		log.Panic(err)
	}

	envReader := env.NewReader(file)
	envReader.Read()
}
