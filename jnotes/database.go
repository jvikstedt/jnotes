package jnotes

type Database interface {
	Setup(databaseURL string) error
	Shutdown()
	CreateTables()
	RecreateTables()
}
