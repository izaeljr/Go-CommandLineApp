package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Start function Generate
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IP's and server names around the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP's addresses on internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, exception := net.LookupIP(host)

	if exception != nil {
		log.Fatal(exception)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
