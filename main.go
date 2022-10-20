package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var directory = flag.String("directory", ".", "Directory")
var address = flag.String("address", ":8080", "IP address and port to listen connections")

func main() {
	flag.Parse()

	fmt.Printf("Listening address=%q\n", *address)
	fmt.Printf("Serving directory=%q\n", *directory)

	dir := http.Dir(*directory)
	handler := http.FileServer(dir)
	http.Handle("/", handler)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
