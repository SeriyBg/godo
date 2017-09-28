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

func Complete(c *cli.Context) (err error) {
	id := c.Args().Get(0)
	if len(id) != 0 {
		println(id)
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "TODO task name",
				},
				cli.StringFlag{
					Name:  "description, d",
					Usage: "TODO task description",
				},
			},
			Action: Add,
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Shows TODO tasks",
			Action:  Show,
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "Complete TODO task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "TODO task name",
				},
			},
			Action: Complete,
		},
	}
	app.Run(os.Args)
}
