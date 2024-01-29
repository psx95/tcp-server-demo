package main

import (
	"fmt"
	"log"
	"net"
	"time"

	netutil "golang.org/x/net/netutil"
)

// Potential enhancements:
// 1. Limit number of threads created by server
// 2. Add thread pool
// 3. Connection timeout
// 4. Configure tcp backlog configuration

func main() {
	log.Println("Starting server")
	// STEP 1: Make the server listen to a particular port
	listener, err := net.Listen("tcp", ":1234") // blocking call
	if err != nil {
		log.Fatalf("Error encountered while attempting to listen: %v", err)
	}
	listener = netutil.LimitListener(listener, 100)
	fmt.Printf("Listener obj :%v\n", listener)
	handleClients(listener)
	fmt.Println("Exiting server")
	defer listener.Close()
}

func handleClients(listener net.Listener) {
	for {
		log.Println("Waiting for Client")
		// STEP 2: Allow server to accept the incoming connections from clients
		conn, err := listener.Accept() // blocking call
		if err != nil {
			log.Fatalf("Unable to accept connection from client: %v", err)
		}
		log.Println("Processing request")
		go doWork(conn)
	}
}

func doWork(con net.Conn) {
	defer func() {
		log.Println("Request completed")
	}()
	// Buffer to store the request
	connBuffer := make([]byte, 2048)

	// STEP 3: Read the request sent by the client to the server
	// Read request made by client over the tcp connection
	bytesRead, err := con.Read(connBuffer) // blocking call
	if err != nil {
		log.Fatalf("Unable to read from connection :%v", err)
	}
	fmt.Printf("Read bytes :%v\n", bytesRead)

	// STEP 4: Process the request
	// Mimic long running process
	time.Sleep(10 * time.Second)

	// STEP 5: Write the response to be read by the client
	//Write some response on the connection
	bytesWritten, err := con.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello from Server\r\n"))
	if err != nil {
		log.Fatalf("Unable to send response to client: %v", err)
	}
	fmt.Printf("Wrote %v bytes to response\n", bytesWritten)

	// STEP 6: Close the connection, indicating to the client that the
	// response from server has been sent
	con.Close()
}
