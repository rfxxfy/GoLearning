package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d * %d = %d\n", i, n, i*n)
	}
}
