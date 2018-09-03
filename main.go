package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/nathanburkett/nathanb-api/app"
	"github.com/nathanburkett/nathanb-api/env"
	"fmt"
	"github.com/nathanburkett/nathanb-api/data"
	"github.com/jmoiron/sqlx"
)

func main() {
	dataSource := data.NewSource(env.Must("DB_DSN")).Connect()
	instance := app.NewInstance()

	instance.SetDataSource(dataSource)

}

func getHostAndPort() string {
	return fmt.Sprintf("%s:%s", env.Must("APP_HOST"), env.Must("APP_PORT"))
}

