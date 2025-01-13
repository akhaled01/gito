package main

import (
	"fmt"
	dns "gito/internal/dns/types"
	"log"
)

func main() {
	data := []byte{
		0x00, 0x1a, // ID
		0x01, 0x00, // Flags
		0x00, 0x01, // QDCOUNT
		0x00, 0x00, // ANCOUNT
		0x00, 0x00, // NSCOUNT
		0x00, 0x00, // ARCOUNT
	}

	// Parse the DNS header
	header, err := dns.DecodeHeader(data)
	if err != nil {
		log.Fatal(err)
	}

	// Print the parsed header
	fmt.Printf("Parsed DNS Header: %+v\n", header)

	flags := header.GetFlags()

	fmt.Printf("Parsed DNS Header flags: %+v\n", flags)
}
