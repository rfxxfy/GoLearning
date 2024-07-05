func intersection(mass []int, mass2 [] int) []int {
	counterFirst := make(map[int]int)
	counterSecond := make(map[int]int)
	for _, m := range mass {
		counterFirst[m]++
	}
	for _, m := range mass2 {
		counterSecond[m]++
	}
	ans := make([]int, 0)
	for key, value := range counterFirst {
		for i := 0; i < min(value, counterSecond[key]); i++ {
			ans = append(ans, key)
		}
	}
	return ans
}
