package main

import "strconv"
import "fmt"

func Reverse(n int) int {
	s := strconv.Itoa(n)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	numToRtn, _ := strconv.Atoi(string(runes))
	return numToRtn
}

func main() {
	var highest int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if i*j == Reverse(i*j) && i*j >= highest {
				highest = i * j
				fmt.Println(i, j)
			}
		}
	}
	fmt.Println("Highest: ", highest)
}
