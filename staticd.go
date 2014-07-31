package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var flport string
var flpath string

func init() {
	flag.StringVar(&flport, "port", "8080", "Listen port")
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
	http.ListenAndServe(":"+flport, nil)
}
