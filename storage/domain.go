package storage

import (
	"encoding/json"
	"fmt"
	"time"
)

type Note struct {
	id          string
	name        string
	description string
	state       Status
	created     time.Time
	updated     time.Time
}

type jsonNote struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	State       Status    `json:"state"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

func (n *Note) UnmarshalJSON(b []byte) error {
	jsonNote := &jsonNote{}
	err := json.Unmarshal(b, &jsonNote)
	if err != nil {
		return err
	}
	n.id = jsonNote.Id
	n.name = jsonNote.Name
	n.description = jsonNote.Description
	n.state = jsonNote.State
	n.created = jsonNote.Created
	n.updated = jsonNote.Updated
	return nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNote{
		Id:          n.id,
		Name:        n.name,
		Description: n.description,
		State:       n.state,
		Created:     n.created,
		Updated:     n.updated,
	})
}

func (n Note) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Description: %s, State: %s, Last change: %s",
		n.id, n.name, n.description, n.state, n.updated.Local().Format(time.RFC822))
}

type Status string

func (n *Note) IsRelevant() bool {
	return n.state == New
}

const (
	New  Status = "New"
	Done Status = "Done"
)

type Repository interface {
	Create(name string, description string) (err error)
	GetAll() (notes []Note, err error)
	CompleteById(id string) (err error)
	FindAllBy(filter NoteFilter) (notes []Note, err error)
	DeleteById(id string) (err error)
}

type NoteFilter func(note Note) bool
