package main

import (
	"flag"
	"net/http"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen")
	folder := flag.String("webroot", ".", "Folder to serve")
	flag.Parse()

	if len(*folder) > 1 && (*folder)[:2] == "~/" {
		usr, _ := user.Current()
		fixedPath := strings.Replace(*folder, "~/", usr.HomeDir+"/", 1)
		folder = &fixedPath
	}

	panic(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*folder))))
}
