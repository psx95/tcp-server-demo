package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// STEP 1: Make the server listen to a particular port
	listener, err := net.Listen("tcp", ":1234") // blocking call
	if err != nil {
		log.Fatalf("Error encountered while attempting to listen: %v", err)
	}
	fmt.Printf("Listener obj :%v\n", listener)
	handleClients(listener)
	fmt.Println("Exiting server")
}

func handleClients(listener net.Listener) {
	for {
		// STEP 2: Allow server to accept the incoming connections from clients
		conn, err := listener.Accept() // blocking call
		if err != nil {
			log.Fatalf("Unable to accept connection from client: %v", err)
		}
		fmt.Printf("Accepted connection from client :%v\n", conn.LocalAddr())

		go doWork(conn)
	}
}

func doWork(con net.Conn) {
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
