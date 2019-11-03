package main

import (
	"context"
	"fmt"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <- c1.Done():
		fmt.Println("f1(): ", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
}

func f2(i int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(i) * time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(30) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func main() {
	f1(10)
	f2(3)
}
