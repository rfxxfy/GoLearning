package main

import (
	"fmt"
	"math"
)

func check(x int) bool {
	var sqrt int = int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrt; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a int
	fmt.Scan(&a)
	for i := 2; i <= a; i++ {
		if check(i) {
			fmt.Printf("%d ", i)
		}
	}
}
