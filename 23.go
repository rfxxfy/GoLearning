func fibonacci(n int) []int {
	if n == 1 {
		return []int{1}
	} else if n == 2 {
		return []int{1, 1}
	}
	ans := make([]int, n)
	ans[0], ans[1] = 1, 1
	for i := 2; i < n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans
}
