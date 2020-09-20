package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
)

func main(){
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMs, MX records and Name Servers"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name: "ns",
			Usage: "Looks up the Name Servers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i:=0; i<len(ns); i++{
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:"ip",
			Usage: "Looks Up the IP addresses for a particular host",
			Flags:myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err !=nil {
					fmt.Println(err)
					return err
				}
				for i:=0; i<len(ip); i++ {
					fmt.Println(ip[i])

				}
				return nil
			},

		},
		{
			Name:"cname",
			Usage: "Looks Up CNAME for a particular Host",
			Flags:myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}
					fmt.Println(cname)
				return nil
			},

		},
		{
			Name:"mx",
			Usage: "Look Up for mx records for a particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))

				if err != nil {
					fmt.Println(err)
					return err
				}
				for i:=0; i<len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil

			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)

	}
}

