package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// client makes a set number of concurrent requests to server running on a given port
// on the local machine

var port = flag.Int("port", 1234, "The port on which server is running")

// Send a high enough number and a server might crash if proper safeguards are not set
var requests = flag.Int("req", 4, "Number of concurrent requests to issue to the server")

const ServerHost = "localhost"

var wg sync.WaitGroup

func main() {
	flag.Parse()
	server_url := fmt.Sprintf("http://%s:%d", ServerHost, *port)
	fmt.Println(server_url)

	for i := 1; i <= *requests; i++ {
		fmt.Printf("issuing request #%v\n", i)
		wg.Add(1)
		go issueSimpleGetRequest(server_url)
	}
	wg.Wait()
}

func issueSimpleGetRequest(serverAddress string) {
	defer wg.Done()
	// Construct HTTP GET request
	req, err := http.NewRequest("GET", serverAddress, nil)
	if err != nil {
		log.Printf("failed to generate request: %v\n", err)
	}
	// Send Request to server
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("unable to send request to server: %v", err)
		return
	}
	defer resp.Body.Close() // Close reader for response

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("unable to read response: %v\n", err)
		return
	}
	fmt.Printf("response received: %v\n", string(body))
}
