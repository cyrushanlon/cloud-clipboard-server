package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/cyrushanlon/cloud-clipboard-server/server"
)

var (
	//holds the username and client information/clipboard
	clients map[string]server.Client
)

func hello(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if ok {

		//try and find a clipboard
		if cli, ok := clients[user]; ok {
			if cli.Password == pass && len(cli.Clipboard) != 0 {
				_, err := io.WriteString(w, string(cli.Clipboard))
				if err != nil {
					log.Println(err)
				}
			}
		}

		// a 200 response without a body is possible and means that there isnt a remote clip
		w.WriteHeader(http.StatusOK)
	} else {
		// a forbidden means the username/password is wrong
		w.WriteHeader(http.StatusForbidden)
	}
}

func main() {
	clients = make(map[string]server.Client)

	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Println(err)
	}
}
