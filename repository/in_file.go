package repository

import (
	"encoding/json"
	"time"
)

var storage []Note = make([]Note, 0, 0)
var lastId int = 1

func AddNote(name string, description string) (err error) {
	now := time.Now()
	note := Note{
		id:          lastId,
		name:        name,
		description: description,
		state:       New,
		created:     now,
		updated:     now,
	}
	lastId += 1
	json, err := json.Marshal(note)
	println(string(json))
	storage = append(storage, note)
	return
}

func ShowAll() (notes []Note, err error) {
	AddNote("Test", "Test test test")
	notes = storage
	return
}
