package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	Handlers "github.com/PanGan21/ethereum-api/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	// Construct the url of the provider to connect
	providerHost := os.Getenv("HOST")
	if len(providerHost) == 0 {
		providerHost = "localhost"
	}

	providerPort := os.Getenv("PORT")
	if len(providerPort) == 0 {
		providerPort = "8545"
	}

	url := fmt.Sprintf("http://%s:%s", providerHost, providerPort)

	// Create a client instance to connect to the provider
	client, err := ethclient.Dial(url)

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// Define a single endpoint
	r.Handle("/api/eth/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))
}
