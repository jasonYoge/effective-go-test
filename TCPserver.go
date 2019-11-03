package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	c, err := l.Accept()
	fmt.Println("Connect accept...")
	if err != nil {
		log.Fatal(err)
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}

}
