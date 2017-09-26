package repository

import (
	"time"
)

type Note struct {
	id          int
	Name        string
	Description string
	State       Status
	Created     time.Time
	Updated     time.Time
}

type Status string

func (s *Status) isRelevant() bool {
	return *s == New || *s == InProgress
}

const (
	New        Status = "New"
	InProgress Status = "In progress"
	Done       Status = "Done"
	Outdated   Status = "Outdated"
)
