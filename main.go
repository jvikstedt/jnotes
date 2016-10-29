package main

import (
	"fmt"
	"github.com/jvikstedt/jnotes/database"
	"github.com/jvikstedt/jnotes/repository"
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

	noteRepository := repository.NoteRepository{DB: database.DB}
	//note := jnotes.Note{Title: "I Love Golang", Body: "It's so good!"}
	//note, err = noteRepository.Create(note)

	//note, _ := noteRepository.FindByID(1)
	//fmt.Println(note)

	//note.Body = "It's awesome"
	//note, err = noteRepository.Update(note)
	//fmt.Println(note)
	//fmt.Println(err)

	//note, err := noteRepository.DeleteByID(1)
	//fmt.Println(note)
	//fmt.Println(err)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! adsdsa")
	})
	e.Run(standard.New(":" + port))
	defer database.Shutdown()
}
