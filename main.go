package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jvikstedt/jnotes/database"
	"github.com/jvikstedt/jnotes/repository"
	"log"
	"net/http"
	"os"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing!\n"))
}

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

	note, err := noteRepository.DeleteByID(1)
	fmt.Println(note)
	fmt.Println(err)

	r := mux.NewRouter()
	r.HandleFunc("/", TestHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))
	defer database.Shutdown()
}
