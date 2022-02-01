// bhg-copy-example.go modified from Black Hat Go > CH2 > echo-server > main.go

// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/echo-server/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// >Potential  go run bhg-echo-server.go

package main

import (
	"io"
	"log"
	"net"
	
)

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])

		// Handle the end of file from server ^]
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		// Handle other errors
		if err != nil {
			log.Println("Unexpected error")
			break
		}

		// If we make it here - we got data!
		log.Printf("Received %d bytes: %s\n", size, string(b))

		c := make([]byte,size)
		for i := 0; i<size-1; i++{
			c[i] = b[size-2-i]
		}
		c[size-1] = b[size-1]

		log.Printf("Writing data backward")

		// Send data via conn.Write.
		if _, err := conn.Write(c[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}	
			
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}