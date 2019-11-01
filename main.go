package main

import (
	"effective-go-test/test"
	"fmt"
)

func main() {
	var a = 10
	var b = 20

	fmt.Println(test.Add(a, b))
}