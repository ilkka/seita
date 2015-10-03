package config

import (
	ccli "github.com/codegangsta/cli"
	"github.com/spf13/viper"

	"github.com/ilkka/seita/cli"
)

// Get config value
func Get(c *ccli.Context) {
	name := c.Args().First()
	cli.Printf("%s\n", viper.GetString(name))
}
