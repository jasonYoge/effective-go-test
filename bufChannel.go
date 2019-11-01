package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter + 5; i++ {
		select {
		case res := <-numbers:
			fmt.Println(res)
		default:
			fmt.Println("Nothing more to be done!")
		}
	}
}
