package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "weblambda"
	app.Usage = "Run JavaScript code on AWS Lambda with HTTP call"

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
			Usage: "Install weblambda function on AWS Lambda",
			Action: func(c *cli.Context) {
				install(c.String("role"), c.String("region"), c.Bool("upgrade"))
			},
		},
		{
			Name:  "server",
			Usage: "Start HTTP server",
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
