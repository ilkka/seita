package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ilkka/seita/command"
	"github.com/ilkka/seita/config"
)

func requireArgs(cmd string, num int) func(*cli.Context) error {
	return func(c *cli.Context) error {
		if len(c.Args()) != num {
			return fmt.Errorf("%s requires %d arguments", cmd, num)
		}
		return nil
	}
}

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
			Before:  requireArgs("get", 1),
		},
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Manipulate configuration",
			Subcommands: []cli.Command{
				{
					Name:    "set",
					Aliases: []string{"s"},
					Usage:   "Set configuration variable value",
					Action:  config.Set,
					Before:  requireArgs("set", 2),
				},
				{
					Name:    "get",
					Aliases: []string{"g"},
					Usage:   "Get configuration variable value",
					Action:  config.Get,
					Before:  requireArgs("get", 1),
				},
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List configuration variables",
					Action:  config.List,
				},
			},
		},
	}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.RunAndExitOnError()
}
