package main

import (
	"fmt"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s\n", xType)

	A := a{100, 200.12, "Struct a"}
	//B := b{1, 2, "struct", -1.2}
	var r reflect.Value
	r = reflect.ValueOf(&A).Elem()

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and Value: %v \n", r.Field(i).Interface())
	}
}
