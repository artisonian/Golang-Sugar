/*
   This is a reference document for Go.sugar
*/
package main

import (
	"flag"
	"fmt"
	"http"
	"os"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("point{x:%v, y:%v}", p.x, p.y)
}

var port = flag.String("http", ":8080", "bind port for http server")

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

func pointHandler(w http.ResponseWriter, r *http.Request) {
	// create a Point
	p := &Point{2, 3}

	// set the content type
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, Espresso!\n%v\n", p)
}

func main() {
	http.HandleFunc("/", homepageHandler)
	http.ListenAndServe(*port, nil)
}

const page = `
<!DOCTYPE html>
<html>
  <head>
    <title>Homepage</title>
  </head>
  <body>
    <h1>Homepage</h1>
  </body>
</html>
`
