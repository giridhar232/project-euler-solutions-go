package main

import "fmt"

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	//fmt.Print(" ", n)
	return true
}

func main() {
	var num uint64
	//num = 600851475143
	num = 600851475143
	for i := 2; uint64(i) < num; i++ {
		if isPrime(i) && num%uint64(i) == 0 {
			fmt.Println("Prime Factor: ", i)
		}
	}
}
