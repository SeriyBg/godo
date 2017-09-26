package main

import (
	"fmt"
	"github.com/SeriyBg/godo/repository"
	"github.com/codegangsta/cli"
	"os"
)

func Add(c *cli.Context) (err error) {
	return repository.AddNote(c.String("name"), c.String("description"))
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
