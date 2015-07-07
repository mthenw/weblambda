package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "weblambda"
	app.Usage = "backendless webhooks server on top of AWS Lambda"

	app.Commands = []cli.Command{
		{
			Name: "install",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "role",
				},
				cli.StringFlag{
					Name:  "region",
					Value: "us-east-1",
					Usage: "AWS region name",
				},
				cli.BoolFlag{
					Name:  "upgrade, u",
					Usage: "Upgrade existing function",
				},
			},
			Usage: "install weblambda function on AWS Lambda",
			Action: func(c *cli.Context) {
				install(c.String("role"), c.String("region"), c.Bool("upgrade"))
			},
		},
		{
			Name:  "server",
			Usage: "start HTTP server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "region",
					Value: "us-east-1",
					Usage: "AWS region name",
				},
				cli.StringFlag{
					Name:  "port",
					Value: "8080",
				},
			},
			Action: func(c *cli.Context) {
				server(c.String("region"), c.String("port"))
			},
		},
	}

	app.Run(os.Args)
}
