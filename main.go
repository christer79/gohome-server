package main

import (
	"fmt"
	"os"

	"github.com/christer79/home-server/disks"
	"github.com/christer79/home-server/ip"
	"github.com/christer79/home-server/web"
)

func main() {
	fmt.Println("Filesystems")
	disks := disks.Client{Filesystems: []string{"/", "/boot", "/home"}}
	disks.Write(os.Stdout)
	ip := ip.Client{}
	ip.Write(os.Stdout)
	web.Client{Disks: disks}.Start("9000")
}
