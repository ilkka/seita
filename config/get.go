package config

import (
	"log"

	"github.com/codegangsta/cli"
)

// Get config value
func Get(c *cli.Context) {
	log.Printf("Get value of %v", c.Args().First())
}
