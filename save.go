package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	s := []byte("Data to write\n")

	f1, err := os.Create("f1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("f2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("f3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f3.Close()
	w := bufio.NewWriter(f3)
	m, _ := w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", m)
	w.Flush()

	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f5, err := os.Create("f5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f5.Close()
	n, err = io.WriteString(f5, string(s))
}
