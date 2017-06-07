package disks

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/christer79/gohome-server/sysinfo"
)

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
	log.Println(diskUsage)
	t.Execute(w, diskUsage)
}

func (c Client) Write(w io.Writer) {
	for _, filesystem := range c.Filesystems {
		usage := sysinfo.DiskUsage(filesystem)
		usage.Write(os.Stdout)
	}
}
