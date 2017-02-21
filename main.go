package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello world!")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Println(err)
	}
}
