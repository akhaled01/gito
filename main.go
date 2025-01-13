package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen on UDP port 8080
	address := "localhost:53"
	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on", address)

	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Print the received message and sender's address
		fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr)

		// Respond with a message
		response := []byte("Message received")
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
