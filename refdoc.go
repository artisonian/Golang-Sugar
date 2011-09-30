package main

import (
	"flag"
	"fmt"
	"http"
	"os"
)

var port = flag.String("http", ":8080", "bind port for http server")

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, Espresso!")
}

func main() {
	http.HandleFunc("/", homepageHandler)
	http.ListenAndServe(*port, nil)
}
