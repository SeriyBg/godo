package main

import (
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func Add(c *cli.Context) (err error) {
	//name := c.String("name")
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "godo"
	app.Usage = "Simple CLI TODO application written in Go"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Create a new TODO task",
			Action:  Add,
		},
	}
	app.Run(os.Args)
}

type Note struct {
	Id          int
	Name        string
	Description string
	State       Status
	Created     time.Time
	Updated     time.Time
}

type Status int8

func (s *Status) isRelevant() bool {
	return *s == New || *s == InProgress
}

const (
	New        Status = 0
	InProgress Status = 1
	Done       Status = 2
	Outdated   Status = 3
)
