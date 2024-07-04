func count(n int, mass []int) int {
	counter := make(map[int]int) 
	for _, m := range mass {
		counter[m]++
	}
	return counter[n]
}
