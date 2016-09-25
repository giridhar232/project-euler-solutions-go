package main

func main() {
	baseNums := []int{1, 2, 3, 4, 5}
	num := 11
	numNotFound := true
	for numNotFound {
		for base := range baseNums {
			// if base < 10 && num%base == 0 {
			// 	//break
			// }
			if base == 5 && num%base == 0 {
				numNotFound = false
				//break
			}
			num++
		}
	}

}
