package config

import (
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

// Set sets a config value to the value of an argument.
func Set(c *cli.Context) {
	name := c.Args().First()
	val := c.Args().Get(1)
	viper.Set(name, val)
	write()
}
