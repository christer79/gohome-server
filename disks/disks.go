package disks

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/christer79/gohome-server/sysinfo"
)

type DiskStatusScaled struct {
	All  string
	Used string
	Free string
	Path string
}

type Client struct {
	Filesystems []string
}

func (c Client) usage() (usage []sysinfo.DiskStatus) {
	for _, filesystem := range c.Filesystems {
		usage = append(usage, sysinfo.DiskUsage(filesystem))
	}
	return
}

//WebHandler handler to render disk usage for web
func (c Client) WebHandler(w http.ResponseWriter, r *http.Request) {
	diskUsage := c.usage()
	t, err := template.ParseFiles("html/disks.html")
	if err != nil {
		log.Println(err)
	}
	var scaledUsage []DiskStatusScaled
	devisor := uint64(1)
	unit := "B"
	for _, usage := range diskUsage {
		if usage.All > 90000000 {
			devisor = sysinfo.MB
			unit = "MB"
		}
		if usage.All > 9000000000 {
			devisor = sysinfo.GB
			unit = "GB"
		}

		scaledUsage = append(scaledUsage, DiskStatusScaled{Path: usage.Path, All: fmt.Sprintf("%2d %s", usage.All/devisor, unit), Free: fmt.Sprintf("%2.2d %s", usage.Free/devisor, unit), Used: fmt.Sprintf("%2.2d %s", usage.Used/devisor, unit)})
	}
	log.Println(scaledUsage)
	t.Execute(w, scaledUsage)
}

func (c Client) Write(w io.Writer) {
	for _, filesystem := range c.Filesystems {
		usage := sysinfo.DiskUsage(filesystem)
		usage.Write(os.Stdout)
	}
}
