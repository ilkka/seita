package config

import (
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"

	c "github.com/ilkka/seita/cli"
)

// List config values
func List(ctx *cli.Context) {
	keys := viper.AllKeys()
	for idx := 0; idx < len(keys); idx++ {
		c.Printf("%s -- %s\n", keys[idx], viper.Get(keys[idx]))
	}
}
