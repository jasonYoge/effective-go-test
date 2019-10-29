package memoryStatus

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name())
	defer os.Remove(f.Name())

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var mem runtime.MemStats
	PrintStatus(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000)
		if s == nil {
			fmt.Println("failed operation!")
		}
	}
	PrintStatus(mem)
}
