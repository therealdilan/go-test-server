package main

import (
	"net/http"

	"github.com/therealdilan/go-test-server/api"
)

func main() {
	newServer := api.NewServer()
	http.ListenAndServe(":8080", newServer)
}
