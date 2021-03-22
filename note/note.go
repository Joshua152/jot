package note

import "time"

var Notes []Note

/* Note represents a single note */
type Note struct {
	Note        string `json:"note"`
	TimeCreated string `json:"timeCreated"`
}

func New(str string) Note {
	return Note{Note: str, TimeCreated: time.Now().Local().String()}
}

func Add(n Note) {
	Notes = append(Notes, n)
}
