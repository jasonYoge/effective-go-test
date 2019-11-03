package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipAddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := net.ListenIP("ip4:icmp", ipAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFrom(buffer)
	fmt.Println("Read from: ", addr.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Received message: ")
	fmt.Printf("% X",  buffer[0:n])
}

