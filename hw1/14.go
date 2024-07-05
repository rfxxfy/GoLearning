func contains(s int, mass []int) bool {
	for _, m := range mass {
		if m == s {
			return true
		}
	}
	return false
}
