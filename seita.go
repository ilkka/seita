package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
  app.Name = "seita"
  app.Usage = "Enshrine and retrieve project skeletons"
  app.Action = func(c *cli.Context) {
    println("Nothing to do!")
  }

  app.Run(os.Args)
}
