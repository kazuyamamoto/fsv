package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var directory = flag.String("directory", ".", "Serving directory")
var address = flag.String("address", ":8080", "Listening address and port")

func main() {
	flag.Parse()

	fmt.Printf("Listening %v\n", *address)

	path, err := filepath.Abs(*directory)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Serving %v\n", path)

	fmt.Println("Ctrl + C to quit.")

	dir := http.Dir(path)
	handler := http.FileServer(dir)
	http.Handle("/", handler)

	if http.ListenAndServe(*address, nil) != nil {
		log.Fatalln(err)
	}
}
