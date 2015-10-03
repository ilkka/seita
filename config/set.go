package config

import (
	"log"

	"github.com/codegangsta/cli"
)

// Set sets a config value to the value of an argument.
func Set(c *cli.Context) {
	log.Printf("Set %v to %v", c.Args().Get(0), c.Args().Get(1))
}
