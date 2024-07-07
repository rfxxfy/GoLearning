func massToSet(mass []int) []int {
	exist := make(map[int]bool)
	var ans []int
	for _, v := range mass {
		if !exist[v] {
			ans = append(ans, v)
			exist[v] = true
		}
	}
	return ans
}
