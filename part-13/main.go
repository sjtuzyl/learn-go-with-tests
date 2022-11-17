package main

import (
	"learn/part-13/server"
	"log"
	"net/http"
)

func main() {
	serv := &server.PlayerServer{Store: server.NewStuPlayerStore()}

	if err := http.ListenAndServe(":5000", serv); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
