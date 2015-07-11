package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen")
	flag.Parse()

	folder := flag.Arg(0)
	if folder == "" {
		folder = "."
	} else if strings.HasPrefix(folder, "~/") {
		usr, _ := user.Current()
		folder = strings.Replace(folder, "~/", usr.HomeDir+"/", 1)
	}

	if _, err := os.Stat(folder); err != nil {
		log.Fatal(err)
	} else {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(folder))))
	}
}
