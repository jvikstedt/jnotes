package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jvikstedt/jnotes/jnotes"
)

type NoteTestData struct {
	Notes []jnotes.Note
}

var noteTestData = NoteTestData{
	Notes: []jnotes.Note{
		{Title: "Golang", Body: "Is Awesome"},
		{Title: "Go", Body: "Is Great"},
		{Title: "Javascript", Body: "Is Not"},
	},
}

func TestCreateNote(t *testing.T) {
	for i, source := range noteTestData.Notes {
		noteAsJSON, _ := json.Marshal(source)
		reader := strings.NewReader(string(noteAsJSON))
		req, _ := http.NewRequest("POST", BaseURL+"/notes", reader)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 201 {
			t.Errorf("Expected status 201 but found: %d", res.StatusCode)
		}

		note := jnotes.Note{}
		json.NewDecoder(res.Body).Decode(&note)

		if note.Title != source.Title {
			t.Error("Expected Golang but found: " + note.Title)
		}

		noteTestData.Notes[i] = note
	}
}

func TestDeleteNote(t *testing.T) {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/notes/%d", BaseURL, noteTestData.Notes[2].ID), nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200 but found: %d", res.StatusCode)
	}

	noteTestData.Notes = append(noteTestData.Notes[:2], noteTestData.Notes[2+1:]...)
}

func TestUpdateNote(t *testing.T) {
	reader := strings.NewReader(`{"title": "Elixir"}`)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("%s/notes/%d", BaseURL, noteTestData.Notes[1].ID), reader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200 but found: %d", res.StatusCode)
	}

	note := jnotes.Note{}
	json.NewDecoder(res.Body).Decode(&note)

	if note.Title != "Elixir" {
		t.Error("Expected Golang but found: " + note.Title)
	}

	noteTestData.Notes[1] = note
}

func TestGetAllNotes(t *testing.T) {
	res, err := http.Get(BaseURL + "/notes")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200 but found: %d", res.StatusCode)
	}

	notes := []jnotes.Note{}
	json.NewDecoder(res.Body).Decode(&notes)

	if len(notes) != len(noteTestData.Notes) {
		t.Errorf("Expected length of: %d but found: %d ", len(noteTestData.Notes), len(notes))
	}

	for i, note := range notes {
		if note.Title != noteTestData.Notes[i].Title {
			t.Errorf("Expected title of %s but found: %s", noteTestData.Notes[i].Title, note.Title)
		}
	}
}

func TestGetNote(t *testing.T) {
	for _, source := range noteTestData.Notes {
		res, err := http.Get(fmt.Sprintf("%s/notes/%d", BaseURL, source.ID))
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("Expected status 200 but found: %d", res.StatusCode)
		}

		note := jnotes.Note{}
		json.NewDecoder(res.Body).Decode(&note)
		if source.Title != note.Title {
			t.Errorf("Expected title of %s but found: %s", source.Title, note.Title)
		}
	}
}
