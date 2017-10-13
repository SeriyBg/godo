package main

import (
	"fmt"
	"github.com/SeriyBg/godo/storage"
	"github.com/codegangsta/cli"
)

func Add(c *cli.Context, repository storage.Repository) (err error) {
	return repository.Create(c.String("name"), c.String("description"))
}

func Show(c *cli.Context, repository storage.Repository) (err error) {
	notes, err := repository.GetAll()
	for _, note := range notes {
		fmt.Println(note)
	}
	return
}

func Complete(c *cli.Context, repository storage.Repository) (err error) {
	id := c.String("id")
	if len(id) != 0 {
		repository.CompleteById(id)
	}
	return
}
