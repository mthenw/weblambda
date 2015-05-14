package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "weblambda"
	app.Usage = "backendless webhooks backed by AWS Lambda"

	app.Commands = []cli.Command{
		{
			Name: "install",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "role",
				},
			},
			Usage: "install weblambda function on AWS Lambda",
			Action: func(c *cli.Context) {
				install(c.String("role"))
			},
		},
	}

	app.Run(os.Args)
}
