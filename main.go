package main

import (
	"fmt"
	"os"

	"github.com/christer79/gohome-server/disks"
	"github.com/christer79/gohome-server/ip"
	"github.com/christer79/gohome-server/web"
	"github.com/urfave/cli/altsrc"
	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Name = "gohome"
	app.Usage = "Home server monitoring."

	app.Flags = []cli.Flag{
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{
			Name:  "disks",
			Usage: "List of disks to monitor usage of",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "port",
			Usage: "Port to listen for incomming requests",
			Value: "9000",
		}),
		cli.StringFlag{
			Name:  "config",
			Value: "config.yml",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(c.GlobalStringSlice("disks"))
		disks := disks.Client{Filesystems: c.StringSlice("disks")}
		ip := ip.Client{}

		web.Client{Disks: disks, IP: ip}.Start(c.String("port"))
		return nil
	}

	app.Before = altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	app.Run(os.Args)
}
