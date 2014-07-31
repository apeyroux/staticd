package main

import (
	"flag"
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
		log.Printf("%s %s %s %s", r.Method, r.RemoteAddr, r.UserAgent(), r.URL.Path)
		http.ServeFile(w, r, flpath+r.URL.Path[1:])
	})
	http.ListenAndServe(":"+flport, nil)
}
