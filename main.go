package main

import (
	"fmt"
	"net/http"

	"github.com/therealdilan/go-test-server/api"
)

func main() {
	newServer := api.NewServer()
	err := http.ListenAndServe(":8080", newServer)
	if err != nil {
		fmt.Println("Server Error: ", err)
	} else {
		fmt.Println("Server Started")
	}
}
