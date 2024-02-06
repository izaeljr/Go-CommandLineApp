package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IP's and server names around the internet"

	return app

}
