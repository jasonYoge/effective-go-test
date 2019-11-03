package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var timeVal = time.Duration(time.Second)

func Timeout(network, host string) (net.Conn, error)  {
	conn, err := net.DialTimeout(network, host, timeVal)
	if err != nil {
		return nil, err
	}
	err = conn.SetDeadline(time.Now().Add(timeVal))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL TIMEOUT\n", filepath.Base(os.Args[0]))
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using Default Timeout!")
		} else {
			timeVal = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	URL := os.Args[1]
	t := http.Transport{
		Dial: Timeout,
	}

	client := http.Client{
		Transport: &t,
	}
	data, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
