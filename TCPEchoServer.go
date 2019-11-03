package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatal("Input length must be 2!")
	}
	SERVER := "localhost" + ":" + arguments[1]
	s, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println("Create TCP Addr error: ", err.Error())
		return
	}
	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}

		fmt.Print("> ", string(buffer[0:n-1]))
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
