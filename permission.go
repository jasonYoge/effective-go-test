package main

import (
	"fmt"
	"os"
)

func main() {
	//arguments := os.Args
	//if len(arguments) == 1 {
	//	log.Fatal("arguments must be 2 unary")
	//}

	filename := "README.md"
	info, _ := os.Stat(filename)
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
}
