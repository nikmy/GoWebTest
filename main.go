package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.FileServer(http.Dir("./static"))
	http.Handle("/", srv)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Unexpected error")
	}
}
