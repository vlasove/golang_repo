package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Print hello world."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Use to say hello to",
		},
		cli.BoolFlag{
			Name:  "spanish, s",
			Usage: "Use Spanish language",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("n")
		if c.GlobalBool("s") {
			fmt.Printf("Hola %s!\n", name)
		} else {
			fmt.Printf("Hello %s!\n", name)

		}

		return nil
	}

	app.Run(os.Args)
}
