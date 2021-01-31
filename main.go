package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "github.com/PanGan21/ethereum-api/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	// Create a client instance to connect to the provider
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// Define a single endpoint
	r.Handle("/api/eth/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))
}
