package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jvikstedt/jnotes/jnotes"
)

type Database struct {
	DB *gorm.DB
}

func (database *Database) Setup(databaseURL string) error {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return err
	}
	database.DB = db
	db.AutoMigrate(&jnotes.Note{})
	return err
}

func (database *Database) Shutdown() {
	database.DB.Close()
}
