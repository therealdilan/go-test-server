package main

import (
	"net/http"

	"./api"
)

func main() {
	newServer := api.NewServer()
	http.ListenAndServe(":8080", newServer)
}
