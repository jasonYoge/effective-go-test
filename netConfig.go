package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		fmt.Println("Interface Flags: ", i.Flags.String())
		fmt.Println("Interface MTU: ", i.MTU)
		fmt.Println("Interface Hardware Address", i.HardwareAddr)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			log.Fatal(err)
		}


		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
