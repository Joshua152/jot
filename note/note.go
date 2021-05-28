package note

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Joshua152/jot/logging"
)

var (
	notePath = `.\note\notes.json`
)

/*Note represents a single note*/
type Note struct {
	Note        string `json:"note"`
	TimeCreated string `json:"timeCreated"`
}

/*New creates a new note*/
func New(str string) Note {
	note := Note{Note: str, TimeCreated: time.Now().Local().String()}

	fmt.Println("Writing to a file... hopefully")

	log.SetOutput(logging.GetLogFile())
	log.SetPrefix("note.go: ")
	log.SetFlags(log.Lshortfile)

	return note
}

/*Add adds a new note to the note list and saves it to the file*/
func Add(n Note) {
	notes, err := GetNotes()
	if err != nil {
		log.Println(err)
	}

	notes = append(notes, n)

	err = saveNotes(notes)
	if err != nil {
		log.Println(err)
	}
}

/*Remove deletes a note from the slice at the specified index*/
func Remove(index int) {
	notes, err := GetNotes()
	if err != nil {
		log.Fatal(err)
	}

	notes = append(notes[:index], notes[index+1:]...)

	err = saveNotes(notes)
	if err != nil {
		log.Println(err)
	}
}

/*GetNotes returns the list of notes*/
func GetNotes() ([]Note, error) {
	f, err := os.OpenFile(notePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var notes []Note

	if len(b) != 0 {
		err = json.Unmarshal(b, &notes)
	}

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func saveNotes(notes []Note) error {
	b, err := json.MarshalIndent(notes, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(notePath, b, 0644) // 0 - file, 1 - x (execute), 2 - w (write), 4 - r (read)
	if err != nil {
		return err
	}

	return nil
}
