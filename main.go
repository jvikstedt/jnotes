package main

import (
	"github.com/jvikstedt/jnotes/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	database := database.Database{}
	err := database.Setup(databaseURL)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! adsdsa")
	})
	e.Run(standard.New(":" + port))
	defer database.Shutdown()
}
