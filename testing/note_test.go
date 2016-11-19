package testing

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/jvikstedt/jnotes/jnotes"
)

func TestCreateNote(t *testing.T) {
	noteJSON := `{"title": "Golang", "body": "Is awesome"}`

	reader := strings.NewReader(noteJSON)
	req, _ := http.NewRequest("POST", BaseURL+"/notes", reader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 201 {
		t.Error("Expected status 201 but found: " + strconv.Itoa(res.StatusCode))
	}
	note := jnotes.Note{}
	json.NewDecoder(res.Body).Decode(&note)

	if note.Title != "Golang" {
		t.Error("Expected Golang but found: " + note.Title)
	}
}
