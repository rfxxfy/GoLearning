func get_max(x int, y int, z int) int {
	var ans int
	switch {
	case x >= y && x >= z:
		ans = x
	case y >= z && y >= x:
		ans = y
	default:
		ans = z
	}
	return ans
}

/*func get_max(x int, y int, z int) int {
	return max(x, y, z)
}*/
