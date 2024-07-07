func remove(mass []int, x int) int {
	ans := -1
	for ind, elem := range mass {
		if elem == x {
			ans = ind
			break
		}
	}
	return ans
}
