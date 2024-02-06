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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP's addresses on internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search server names on internet",
			Flags:  flags,
			Action: searchServers,
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

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, exception := net.LookupNS(host)

	if exception != nil {
		log.Fatal(exception)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
