package main

import (
	"fmt"
)

func main() {
	// var xs []int
	xs := []int{}
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			xs = append(xs, i)
		}
	}
	var total int
	for _, v := range xs {
		total += v
	}
	fmt.Println(total)
}
