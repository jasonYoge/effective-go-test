package main

import "fmt"

func main() {
	aMap := map[string]int{}

	aMap["test"] = 1

	fmt.Println(aMap)

	fmt.Println(len(aMap))

	aMap = make(map[string]int, 5)

	fmt.Println(aMap)
}