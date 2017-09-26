package repository

import "time"

var storage []Note = make([]Note, 0, 0)
var lastId int = 1

func AddNote(note Note) (err error) {
	note.id = lastId
	lastId += 1
	storage = append(storage, note)
	return
}

func ShowAll() (notes []Note, err error) {
	AddNote(Note{
		Name:        "Test",
		Description: "Test test test",
		Created:     time.Now(),
		Updated:     time.Now(),
		State:       New,
	})
	notes = storage
	return
}
