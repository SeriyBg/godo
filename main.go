package main

import (
	"fmt"
	"github.com/SeriyBg/godo/storage"
	"github.com/codegangsta/cli"
	"os"
)

var repository = storage.NewRepository()

func Add(c *cli.Context) (err error) {
	return repository.Create(c.String("name"), c.String("description"))
}

func Show(c *cli.Context) (err error) {
	notes, err := repository.GetAll()
	for _, note := range notes {
		fmt.Println(note)
	}
	return
}

func Complete(c *cli.Context) (err error) {
	id := c.String("id")
	if len(id) != 0 {
		repository.CompleteById(id)
	}
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "godo"
	app.Usage = "Simple CLI To-Do application written in Go"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Create a new To-Do task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "To-Do task name",
				},
				cli.StringFlag{
					Name:  "description, d",
					Usage: "To-Do task description",
				},
			},
			Action: Add,
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Shows To-Do tasks",
			Action:  Show,
			Flags:   []cli.Flag{},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "Complete To-Do task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id, i",
					Usage: "To-Do task id",
				},
			},
			Action: Complete,
		},
	}
	app.Run(os.Args)
}
