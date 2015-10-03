package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/ilkka/seita/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "seita"
	app.Usage = "Enshrine and retrieve project skeletons"
	app.EnableBashCompletion = true
	app.Authors = []cli.Author{
		{
			Name:  "Ilkka Laukkanen",
			Email: "ilkka@ilkka.io",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "put",
			Aliases: []string{"p"},
			Usage:   "Offer up this project as a skeleton",
			Action:  command.Put,
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get a skeleton for a new project",
			Action:  command.Get,
		},
	}

	app.Action = func(c *cli.Context) {
		println("Nothing to do!")
	}

	app.Run(os.Args)
}
