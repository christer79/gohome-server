package main

import (
	"fmt"
	"os"

	"github.com/christer79/gohome-server/disks"
	"github.com/christer79/gohome-server/ip"
	"github.com/christer79/gohome-server/web"
)

func main() {
	fmt.Println("Filesystems")
	disks := disks.Client{Filesystems: []string{"/", "/mnt/3000GB", "/mnt/25000GB","/mnt/BACKUP_OTHER","/mnt/BACKUP_PHOTOS"}}
	disks.Write(os.Stdout)
	ip := ip.Client{}
	ip.Write(os.Stdout)
	web.Client{Disks: disks}.Start("9000")
}
