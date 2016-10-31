package main

import (
	"github.com/jvikstedt/jnotes/controller"
	"github.com/jvikstedt/jnotes/database"
	"github.com/jvikstedt/jnotes/repository"
	"github.com/jvikstedt/jnotes/router"
	"log"
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
	defer database.Shutdown()

	noteRepository := repository.NoteRepository{DB: database.DB}

	noteController := controller.NoteController{NoteRepository: noteRepository}

	router := router.Router{NoteController: noteController}

	log.Fatal(http.ListenAndServe(":"+port, router.Handler()))
}
