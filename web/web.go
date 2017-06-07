package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/christer79/gohome-server/disks"
	"github.com/christer79/gohome-server/ip"
	"github.com/christer79/gohome-server/sysinfo"
	"github.com/gorilla/mux"
)

type Client struct {
	Disks disks.Client
	IP    ip.Client
}

//WebHandler handler to render home screen
func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/home.html")
	if err != nil {
		log.Println(err)
	}
	i := 0
	t.Execute(w, i)
}

//WebHandler handler to render home screen
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	sysinfo.Shutdown()
}

//Start start a http server to expose configuration adn status
func (c Client) Start(port string) {
	log.Println("Starting web interface on port: ", port)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/disks/", c.Disks.WebHandler)
	router.HandleFunc("/ip/", c.IP.WebHandler)
	router.HandleFunc("/home/", homeHandler)
	router.HandleFunc("/shutdown/", shutdownHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
