package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	pgport = "PG_PORT"
	pghost = "PG_HOST"
	pguser = "PG_USER"
	pgpass = "PG_PASS"
	pgname = "PG_NAME"
)

func (app App) Initialize() {
	config := dbConfig()

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[pghost], config[pgport],
		config[pguser], config[pgpass], config[pgname])

	app.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = app.DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(pghost)
	if !ok {
		panic("PG_HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(pgport)
	if !ok {
		panic("PG_PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(pguser)
	if !ok {
		panic("PG_USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(pgpass)
	if !ok {
		panic("PG_PASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(pgname)
	if !ok {
		panic("PG_NAME environment variable required but not set")
	}
	conf[pghost] = host
	conf[pgport] = port
	conf[pguser] = user
	conf[pgpass] = password
	conf[pgname] = name

	return conf
}
