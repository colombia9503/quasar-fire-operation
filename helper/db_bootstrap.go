package helper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var dbConnectionInstance = new(dbConnection)

type dbConnection struct {
	Db *sql.DB
}

func (dbc *dbConnection) BootstrapDBConnection(user, password, dbname, host string, port int) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	var err error
	if dbc.Db, err = sql.Open("postgres", connStr); err != nil {
		log.Fatal("Error connecting to database", err)
	}

	dbc.Db.SetMaxIdleConns(10)
	dbc.Db.SetMaxOpenConns(100)
	dbc.Db.SetConnMaxLifetime(time.Hour)

	if err = dbc.Db.Ping(); err != nil {
		log.Fatal("Error pinging to database", err)
	}
}
