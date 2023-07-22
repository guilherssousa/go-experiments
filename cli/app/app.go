package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generates the base CLI app, ready to be executed.
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca endereços IP na internet usando Hosts",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "host",
			Usage:  "Busca endereços na internet usando IPs",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
