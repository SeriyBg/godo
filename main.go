package main

import (
	"fmt"
	"github.com/SeriyBg/godo/repository"
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func Add(c *cli.Context) (err error) {
	now := time.Now()
	note := repository.Note{
		Name:        c.String("name"),
		Description: c.String("description"),
		State:       repository.New,
		Created:     now,
		Updated:     now,
	}
	return repository.AddNote(note)
}

func Show(c *cli.Context) (err error) {
	notes, err := repository.ShowAll()
	for _, note := range notes {
		fmt.Println(note)
	}
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
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Shows TODO tasks",
			Action:  Show,
		},
	}
	app.Run(os.Args)
}
