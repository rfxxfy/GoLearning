func sum(x []int) int {
	res := 0
	for i := 0; i < len(x); i++ {
		res += x[i]
	}
	return res
}
