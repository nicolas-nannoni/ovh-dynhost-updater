package main

import (
	"./config"
	"./ovh"

	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// App build information
var (
	Version string
	Build   string
)

func initApp(c *cli.Context) error {

	if config.Config.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled!")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	return nil
}

func appDefinition() (app *cli.App) {

	app = cli.NewApp()
	app.Name = "ovh-dynhost-upgrader"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Nicolas Nannoni",
			Email: "nannoni@kth.se",
		},
	}
	app.Usage = "Tool to perform DynHost DNS records for OVH domains."
	app.Before = initApp
	app.Version = Version + " (build " + Build + ")"

	app.Commands = []cli.Command{
		{
			Name:   "update-record",
			Usage:  "Update a DynHost record",
			Action: ovh.UpdateRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "ip-address",
					Value:       "",
					Usage:       "The IP address that should be used to update the DynHost (bypass auto-detection).",
					Destination: &config.Config.IpAddress,
				},
				cli.StringFlag{
					Name:        "interface, I",
					Value:       "",
					Usage:       "The interface whose IP address should be used to update the DynHost.",
					Destination: &config.Config.NetworkInterface,
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "username",
			Usage:       "The OVH DynHost username",
			Destination: &config.Config.Username,
		},
		cli.StringFlag{
			Name:        "password",
			Usage:       "The OVH DynHost password",
			Destination: &config.Config.Password,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug mode",
			Destination: &config.Config.Debug,
		},
	}

	return
}

func main() {
	appDefinition().Run(os.Args)
}
