package ip

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/christer79/home-server/sysinfo"
)

//Client ip address client
type Client struct {
}

//WebHandler handler to render disk usage for web
func (c Client) WebHandler(w http.ResponseWriter, r *http.Request) {
	ip := c.getIP()
	t, err := template.ParseFiles("html/ip.html")
	if err != nil {
		log.Println(err)
	}
	log.Println(ip)
	t.Execute(w, ip)
}

func (c Client) getIP() (ips []net.IP) {
	ips = sysinfo.IPAdresses()
	return
}

func (c Client) Write(w io.Writer) {
	for _, ip := range c.getIP() {
		fmt.Fprintf(w, "IP: %s\n", ip.String())
	}
}
