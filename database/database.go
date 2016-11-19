package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS notes (
	id serial NOT NULL,
	title character varying(100) NOT NULL,
	body text,
	createdAt timestamptz,
	updatedAt timestamptz default current_timestamp,
	CONSTRAINT notes_pkey PRIMARY KEY (id)
);`

type Database struct {
	DB *sqlx.DB
}

func (database *Database) Setup(databaseURL string) error {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return err
	}
	database.DB = db

	database.CreateTables()
	return err
}

func (database *Database) Shutdown() {
	database.DB.Close()
}

func (database *Database) CreateTables() {
	database.DB.MustExec(schema)
}

func (database *Database) RecreateTables() {
	database.DB.MustExec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`)
	database.CreateTables()
}
