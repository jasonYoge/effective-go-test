package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	if c := recover(); c != nil {
		fmt.Println("before a()")
	}
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About call to b()")
	b()
	fmt.Println("b() excuted!")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")
}
