package main

import "fmt"

func factorial(x int) int {
	ans := 1
	for i := 2; i <= x; i++ {
		ans *= i
	}
	return ans
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(factorial(a))
}
