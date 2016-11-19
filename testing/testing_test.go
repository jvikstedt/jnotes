package testing

import (
	"flag"
	"os"
	"testing"

	"github.com/jvikstedt/jnotes/database"
)

var BaseURL = "http://nginx/api/v1"

type Helper struct {
	DB database.Database
}

var helper Helper

func TestMain(m *testing.M) {
	flag.Parse()
	SetupDB()
	os.Exit(m.Run())
}

func SetupDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	helper = Helper{}
	helper.DB = database.Database{}
	err := helper.DB.Setup(databaseURL)
	if err != nil {
		panic(err)
	}
	helper.DB.RecreateTables()
}
