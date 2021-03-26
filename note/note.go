package note

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

/*Note represents a single note*/
type Note struct {
	Note        string `json:"note"`
	TimeCreated string `json:"timeCreated"`
}

/*New creates a new note*/
func New(str string) Note {
	note := Note{Note: str, TimeCreated: time.Now().Local().String()}

	log.SetFlags(log.Lshortfile)

	return note
}

/*Add addds a new note to the note list and saves it to the file*/
func Add(n Note) {
	notes, err := getNotes()
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	notes = append(notes, n)

	err = saveNotes(notes)
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

}

func saveNotes(notes []Note) error {
	list, err := getNotes()
	if err != nil {
		return err
	}

	list = append(notes)

	b, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("note/notes.json", b, 0644) // 0 - file, 1 - x (execute), 2 - w (write), 4 - r (read)
	if err != nil {
		return err
	}

	return nil
}

func getNotes() ([]Note, error) {
	b, err := ioutil.ReadFile("note/notes.json")
	if err != nil {
		fmt.Println("error1")
		return nil, err
	}

	var notes []Note

	if len(b) != 0 {
		err = json.Unmarshal(b, &notes)
	}

	if err != nil {
		fmt.Println("error2")
		return nil, err
	}

	return notes, nil
}
