package main

import "fmt"

func check(x int) bool {
	return x%2 == 0
}

func main() {
	var a int
	fmt.Scan(&a)
	if check(a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
