package main

import (
	"fmt"
)

func sum_natrl_nums(l int, h int) {
	var total int
	for i := l; i < h; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Println("Total of all multiples of 3 or 5 till ", h, " is : ", total)
}

func main() {
	sum_natrl_nums(1, 1000)
}
