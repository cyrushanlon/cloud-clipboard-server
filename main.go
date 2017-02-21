package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello world!")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":443", nil)
	if err != nil {
		log.Println(err)
	}
}
