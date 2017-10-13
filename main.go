package main

import (
	"github.com/SeriyBg/godo/storage"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	repository := storage.NewRepository()
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
			Action: getCliAction(repository, Add),
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Shows To-Do tasks",
			Action:  getCliAction(repository, Show),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "complete, c",
					Usage: "Show complete To-Do tasks",
				},
				cli.StringFlag{
					Name:  "relevant, r",
					Usage: "Show relevant To-Do tasks",
				},
			},
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
			Action: getCliAction(repository, Complete),
		},
	}
	app.Run(os.Args)
}

type repoAction func(c *cli.Context, repo storage.Repository) error

func getCliAction(repo storage.Repository, action repoAction) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		return action(c, repo)
	}
}
