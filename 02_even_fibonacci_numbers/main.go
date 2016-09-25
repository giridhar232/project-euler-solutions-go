package main

import "fmt"

func even_fibonacci(n int) {
	num1 := 1
	num2 := 2
	var temp, total int
	fmt.Print(num1, " ")
	for num2 < n {
		fmt.Print(num2, " ")
		if num2%2 == 0 {
			total += num2
		}
		temp = num1
		num1 = num2
		num2 = temp + num1
	}
	fmt.Println("\nTotal: ", total)
}

func main() {
	even_fibonacci(4000000)
}
