package main

import (
	"fmt"
)

func main() {
	vslice := []int{1, 2, 24, 5}
	fmt.Println(getMax(vslice))
}

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
}
