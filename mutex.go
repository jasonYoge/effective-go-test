package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM sync.RWMutex
	M sync.Mutex
	password string
}

func change(c *secret, pass string) {
	c.RWM.Lock()
	defer c.RWM.Unlock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
}

func show(c *secret) string {
	c.RWM.RLock()
	defer c.RWM.RUnlock()
	fmt.Println("show")
	time.Sleep(3 * time.Second)
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	defer c.M.Unlock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	return c.password
}

var Password = secret{ password: "myPassword" }

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}

	var waitGroup sync.WaitGroup
	fmt.Println("Pass: ", showFunction(&Password))

	for i := 0 ; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()

		go func(){
			waitGroup.Add(1)
			defer waitGroup.Done()
			change(&Password, "123456")
		}()
		waitGroup.Wait()
		fmt.Println("Pass:", showFunction(&Password))
	}
}
