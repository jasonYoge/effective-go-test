package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name string
	SurName string
	Id string
}

var DATA = make(map[string]myElement)

func LOOKUP(k string) *myElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}

	return false
}

func ADD(k string, n myElement) bool {
	if k == "" {
		return false
	}
	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}

	return false
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("key: %s value: %v",k,v)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

	}
}
