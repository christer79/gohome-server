package sysinfo

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"syscall"
)

//DiskStatus struct to keep disk path and usage
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Path string `json:"path"`
}

func (d DiskStatus) Write(w io.Writer) {

	var devisor uint64
	devisor = GB
	unit := "GB"
	if d.All < 900000000 {
		devisor = MB
		unit = "MB"
	}

	fmt.Fprintf(w, "Folder: %s \n  Free: %v %s Used: %v %s Total: %v %s \n", d.Path, d.Free/devisor, unit, d.Used/devisor, unit, d.All/devisor, unit)
}

func Shutdown() (err error) {

	cmd := exec.Command("shutdown", "now")
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		log.Println(err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			log.Printf("%d\n", waitStatus.ExitStatus())
		}
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("%d\n", waitStatus.ExitStatus())
	}

	cmd.Run()
	return

}

//IPAdresses returns a list of configured ip adresses
func IPAdresses() (ip []net.IP) {
	ifaces, _ := net.Interfaces()
	// TODO: handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// TODO: handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = append(ip, v.IP)
			case *net.IPAddr:
				ip = append(ip, v.IP)
			}
		}
	}
	return
}

//DiskUsage disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		log.Printf("Error in Statfs: %v\n", err)
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	disk.Path = path
	return
}

//Constants for B,KB,MB,GB conversion
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)
