package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	width int
	height int
}

func main() {
	mySlice := []aStructure{
		{
			"ywj1",
			5,
			5,
		},
		{
			"ywj2",
			4,
			4,
		},
		{
			"ywj3",
			3,
			3,
		},
	}

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

	fmt.Println("<:", mySlice)
}
