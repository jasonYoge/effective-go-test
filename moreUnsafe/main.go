package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := []int{0, 1, -2, 3, 4}
	pointer := &arr[0]

	fmt.Println(*pointer, "    ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	for i := 0; i < len(arr) - 1; i++ {
		pointer := (*int)(unsafe.Pointer(memoryAddress))
		fmt.Println(*pointer, "  ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	}

	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One More: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(arr[0])
	fmt.Println()
}

