package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var fllisten string
var flpath string

func init() {
	flag.StringVar(&fllisten, "listen", "127.0.0.1:8080", "Listen port")
	flag.StringVar(&flpath, "path", "", "Root Path")
}

func main() {
	flag.Parse()
	if flpath == "" {
		flag.Usage()
		os.Exit(0)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s", r.Method, r.RemoteAddr, r.UserAgent(), flpath+r.URL.Path[1:])
		url := fmt.Sprintf("%s%s", flpath, r.URL.Path[1:])
		http.ServeFile(w, r, url)
	})
	http.ListenAndServe(fllisten, nil)
}
