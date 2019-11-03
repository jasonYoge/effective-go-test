package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	}
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	c := &http.Client{
		Timeout:       15 * time.Second,
	}

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	httpData, err := c.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status code: ", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
